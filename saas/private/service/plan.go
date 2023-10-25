package service

import (
	"context"
	"fmt"
	"github.com/dtm-labs/dtm/client/dtmcli"
	"github.com/dtm-labs/dtm/client/workflow"
	"github.com/go-kratos/kratos/v2/errors"
	klog "github.com/go-kratos/kratos/v2/log"
	v12 "github.com/go-saas/kit/dtm/api/dtm/v1"
	dtmsrv "github.com/go-saas/kit/dtm/service"
	"github.com/go-saas/kit/pkg/authz/authz"
	"github.com/go-saas/kit/pkg/query"
	"github.com/go-saas/kit/pkg/utils"
	productapi "github.com/go-saas/kit/product/api"
	v1 "github.com/go-saas/kit/product/api/product/v1"
	productbiz "github.com/go-saas/kit/product/private/biz"
	"github.com/go-saas/kit/saas/api"
	"github.com/go-saas/kit/saas/private/biz"
	"github.com/samber/lo"
	"github.com/segmentio/ksuid"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"strings"

	pb "github.com/go-saas/kit/saas/api/plan/v1"
)

var wfCreatePlanName = "saas_create_plan"

type PlanService struct {
	pb.UnimplementedPlanServiceServer
	auth       authz.Service
	repo       biz.PlanRepo
	logger     *klog.Helper
	txhelper   *dtmsrv.Helper
	productSrv v1.ProductInternalServiceServer
}

func NewPlanService(auth authz.Service, repo biz.PlanRepo, logger klog.Logger, txhelper *dtmsrv.Helper, productSrv v1.ProductInternalServiceServer) *PlanService {
	s := &PlanService{auth: auth, repo: repo, logger: klog.NewHelper(klog.With(logger, "module", "PlanService")), txhelper: txhelper, productSrv: productSrv}

	err := s.txhelper.WorkflowRegister2(wfCreatePlanName, func(wf *workflow.Workflow, data []byte) ([]byte, error) {

		//create product ->create plan ->update plan ->delete product if failed
		var req = &pb.CreatePlanRequest{}
		utils.PbMustUnMarshalJson(data, req)

		var productId string
		wf.NewBranch().OnRollback(func(_ *dtmcli.BranchBarrier) error {
			//delete product
			_, err := s.productSrv.DeleteInternalProduct(wf.Context, &v1.DeleteInternalProductRequest{Id: productId})
			return err
		})
		product, err := s.productSrv.CreateInternalProduct(wf.Context,
			&v1.CreateInternalProductRequest{
				Title:      req.DisplayName,
				ManageInfo: &v1.ProductManageInfo{Managed: true, ManagedBy: string(productbiz.ProductManageProviderInternal)},
			})
		if err != nil {
			return nil, err
		}
		productId = product.Id

		resp, err := wf.NewBranch().OnRollback(func(_ *dtmcli.BranchBarrier) error {
			//do nothing
			return nil
		}).Do(func(bb *dtmcli.BranchBarrier) ([]byte, error) {
			resp := &pb.Plan{}
			err := s.txhelper.BarrierUow(wf.Context, bb, biz.ConnName, func(ctx context.Context) error {
				//check duplicate name
				if dbP, err := s.repo.Get(ctx, normalizeName(req.Key)); err != nil {
					return err
				} else if dbP != nil {
					return pb.ErrorDuplicatePlanKeyLocalized(ctx, nil, nil)
				}
				e := &biz.Plan{}
				MapCreatePbPlan2Biz(req, e)
				err := s.repo.Create(ctx, e)
				if err != nil {
					return err
				}
				MapBizPlan2Pb(e, resp)
				return nil
			})
			if err != nil {
				return nil, err
			}
			return utils.PbMustMarshalJson(resp), err
		})
		if err != nil {
			return nil, fmt.Errorf("%s %w", err.Error(), dtmcli.ErrFailure)
		}
		return resp, nil
	})
	if err != nil {
		panic(err)
	}
	return s
}

func (s *PlanService) ListPlan(ctx context.Context, req *pb.ListPlanRequest) (*pb.ListPlanReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourcePlan, "*"), authz.ReadAction); err != nil {
		return nil, err
	}
	ret := &pb.ListPlanReply{}

	totalCount, filterCount, err := s.repo.Count(ctx, req)
	ret.TotalSize = int32(totalCount)
	ret.FilterSize = int32(filterCount)

	if err != nil {
		return ret, err
	}
	items, err := s.repo.List(ctx, req)
	if err != nil {
		return ret, err
	}
	rItems := lo.Map(items, func(g *biz.Plan, _ int) *pb.Plan {
		b := &pb.Plan{}
		MapBizPlan2Pb(g, b)
		return b
	})

	ret.Items = rItems
	return ret, nil
}

func (s *PlanService) GetPlan(ctx context.Context, req *pb.GetPlanRequest) (*pb.Plan, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourcePlan, req.Key), authz.ReadAction); err != nil {
		return nil, err
	}
	g, err := s.repo.Get(ctx, req.GetKey())
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}
	res := &pb.Plan{}
	MapBizPlan2Pb(g, res)
	return res, nil
}

func (s *PlanService) CreatePlan(ctx context.Context, req *pb.CreatePlanRequest) (*pb.Plan, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourcePlan, "*"), authz.CreateAction); err != nil {
		return nil, err
	}

	var err error
	var resp = &pb.Plan{}
	//Workflow Transaction
	data, err := workflow.ExecuteCtx(ctx, wfCreateTenantName, ksuid.New().String(), utils.PbMustMarshalJson(req))
	if err != nil {
		return nil, err
	}
	utils.PbMustUnMarshalJson(data, resp)
	return resp, err
}

func (s *PlanService) UpdatePlan(ctx context.Context, req *pb.UpdatePlanRequest) (*pb.Plan, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourcePlan, req.Plan.Key), authz.UpdateAction); err != nil {
		return nil, err
	}

	//check duplicate name
	if dbP, err := s.repo.Get(ctx, normalizeName(req.Plan.Key)); err != nil {
		return nil, err
	} else if dbP != nil && dbP.Key != req.Plan.Key {
		return nil, pb.ErrorDuplicatePlanKeyLocalized(ctx, nil, nil)
	}
	g, err := s.repo.Get(ctx, req.Plan.Key)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}

	//update plan ->update product(2-phase msg)
	updateReq := &v1.UpdateInternalProductRequest{
		Product:    &v1.UpdateProduct{Id: g.ProductId, Title: g.DisplayName},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"title"}},
	}
	msg := s.txhelper.NewMsgGrpc(ksuid.New().String()).
		Add(productapi.ServiceName+v1.ProductInternalService_UpdateInternalProduct_FullMethodName, updateReq)

	res := &pb.Plan{}
	err = msg.DoAndSubmit(api.ServiceName+v12.MsgService_QueryPrepared_FullMethodName, func(bb *dtmcli.BranchBarrier) error {
		return s.txhelper.BarrierUow(msg.Context, bb, biz.ConnName, func(ctx context.Context) error {
			MapUpdatePbPlan2Biz(req.Plan, g)
			if err := s.repo.Update(ctx, req.Plan.Key, g, nil); err != nil {
				return err
			}
			MapBizPlan2Pb(g, res)
			return nil
		})
	})
	return res, err
}

func (s *PlanService) DeletePlan(ctx context.Context, req *pb.DeletePlanRequest) (*pb.DeletePlanReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourcePlan, req.Key), authz.DeleteAction); err != nil {
		return nil, err
	}

	g, err := s.repo.Get(ctx, req.Key)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}
	//delete plan-> delete product

	msgReq := &v1.DeleteInternalProductRequest{Id: g.ProductId}
	msg := s.txhelper.NewMsgGrpc(ksuid.New().String()).
		Add(productapi.ServiceName+v1.ProductInternalService_DeleteInternalProduct_FullMethodName, msgReq)

	res := &pb.DeletePlanReply{}
	err = msg.DoAndSubmit(api.ServiceName+v12.MsgService_QueryPrepared_FullMethodName, func(bb *dtmcli.BranchBarrier) error {
		return s.txhelper.BarrierUow(msg.Context, bb, biz.ConnName, func(ctx context.Context) error {
			return s.repo.Delete(ctx, req.Key)
		})
	})

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PlanService) GetAvailablePlans(ctx context.Context, req *pb.GetAvailablePlansRequest) (*pb.GetAvailablePlansReply, error) {
	items, err := s.repo.List(ctx, &pb.ListPlanRequest{
		PageOffset: 0,
		PageSize:   -1,
		Filter:     &pb.PlanFilter{Active: &query.BooleanFilterOperators{Eq: &wrapperspb.BoolValue{Value: true}}},
	})
	if err != nil {
		return nil, err
	}
	var retItems = lo.Map(items, func(a *biz.Plan, _ int) *pb.Plan {
		ret := &pb.Plan{}
		MapBizPlan2Pb(a, ret)
		return ret
	})
	return &pb.GetAvailablePlansReply{Items: retItems}, nil
}

func MapBizPlan2Pb(a *biz.Plan, b *pb.Plan) {
	b.Key = a.Key
	b.DisplayName = a.DisplayName
	b.Active = a.Active
}

func MapUpdatePbPlan2Biz(a *pb.UpdatePlan, b *biz.Plan) {

	b.DisplayName = a.DisplayName
	b.Active = a.Active
}

func MapCreatePbPlan2Biz(a *pb.CreatePlanRequest, b *biz.Plan) {
	b.Key = a.Key
	b.DisplayName = a.DisplayName
}

func normalizeName(name string) string {
	return strings.ToLower(name)
}
