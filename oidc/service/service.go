package service

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/go-saas/kit/oidc/api/client/v1"
	v12 "github.com/go-saas/kit/oidc/api/key/v1"
	kitdi "github.com/go-saas/kit/pkg/di"
	"github.com/go-saas/kit/pkg/server"
	"github.com/goava/di"
	client "github.com/ory/hydra-client-go"
	"net/http"
)

var ProviderSet = kitdi.NewSet(
	NewHttpServerRegister,
	NewGrpcServerRegister,
	kitdi.NewProvider(NewClientService, di.As(new(v1.ClientServiceServer))),
	kitdi.NewProvider(NewKeyService, di.As(new(v12.KeyServiceServer))),
)

func NewHttpServerRegister(
	client *ClientService,
	key *KeyService,
) server.HttpServiceRegister {
	return server.HttpServiceRegisterFunc(func(srv *khttp.Server, middleware ...middleware.Middleware) {
		v1.RegisterClientServiceHTTPServer(srv, client)
		v12.RegisterKeyServiceHTTPServer(srv, key)
	})
}

func NewGrpcServerRegister(client *ClientService,
	key *KeyService) server.GrpcServiceRegister {
	return server.GrpcServiceRegisterFunc(func(srv *grpc.Server, middleware ...middleware.Middleware) {
		v1.RegisterClientServiceServer(srv, client)
		v12.RegisterKeyServiceServer(srv, key)
	})
}

func transformErr(resp *http.Response, err error) error {
	if apiError, ok := err.(*client.GenericOpenAPIError); ok {
		if jsonErr, ok := apiError.Model().(client.JsonError); ok {
			reason := ""
			msg := ""
			if jsonErr.Error != nil {
				reason = *jsonErr.Error
			}
			if jsonErr.ErrorDescription != nil {
				msg = *jsonErr.ErrorDescription
			}
			return errors.New(resp.StatusCode, reason, msg)
		}
	}
	return err
}
