// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/goxiaoy/go-eventbus"
	"github.com/goxiaoy/go-saas-kit/pkg/api"
	"github.com/goxiaoy/go-saas-kit/pkg/authn/jwt"
	"github.com/goxiaoy/go-saas-kit/pkg/authz/authz"
	"github.com/goxiaoy/go-saas-kit/pkg/conf"
	"github.com/goxiaoy/go-saas-kit/pkg/dal"
	"github.com/goxiaoy/go-saas-kit/pkg/gorm"
	"github.com/goxiaoy/go-saas-kit/pkg/server"
	"github.com/goxiaoy/go-saas-kit/pkg/uow"
	"github.com/goxiaoy/go-saas-kit/saas/private/biz"
	conf2 "github.com/goxiaoy/go-saas-kit/saas/private/conf"
	"github.com/goxiaoy/go-saas-kit/saas/private/data"
	server2 "github.com/goxiaoy/go-saas-kit/saas/private/server"
	"github.com/goxiaoy/go-saas-kit/saas/private/service"
	api2 "github.com/goxiaoy/go-saas-kit/user/api"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(services *conf.Services, security *conf.Security, confData *conf.Data, saasConf *conf2.SaasConf, logger log.Logger, appConfig *conf.AppConfig, arg ...grpc.ClientOption) (*kratos.App, func(), error) {
	tokenizerConfig := jwt.NewTokenizerConfig(security)
	tokenizer := jwt.NewTokenizer(tokenizerConfig)
	eventBus := _wireEventBusValue
	dbCache, cleanup := gorm.NewDbCache(confData)
	constConnStrResolver := dal.NewConstantConnStrResolver(confData)
	constDbProvider := dal.NewConstDbProvider(dbCache, constConnStrResolver, confData)
	dataData, cleanup2, err := data.NewData(confData, constDbProvider, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	tenantRepo := data.NewTenantRepo(eventBus, dataData)
	tenantStore := data.NewTenantStore(tenantRepo)
	config := _wireConfigValue
	manager := uow.NewUowManager(config, dbCache)
	webMultiTenancyOption := server.NewWebMultiTenancyOption(appConfig)
	option := api.NewDefaultOption(logger)
	decodeRequestFunc := _wireDecodeRequestFuncValue
	encodeResponseFunc := _wireEncodeResponseFuncValue
	encodeErrorFunc := _wireEncodeErrorFuncValue
	trustedContextValidator := api.NewClientTrustedContextValidator()
	clientName := _wireClientNameValue
	inMemoryTokenManager := api.NewInMemoryTokenManager(tokenizer, logger)
	grpcConn, cleanup3 := api2.NewGrpcConn(clientName, services, option, inMemoryTokenManager, logger, arg...)
	userServiceServer := api2.NewUserGrpcClient(grpcConn)
	userTenantContributor := api2.NewUserTenantContributor(userServiceServer)
	connStrGenerator := biz.NewConfigConnStrGenerator(saasConf)
	connName := _wireConnNameValue
	sender, cleanup4, err := dal.NewEventSender(confData, logger, connName)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tenantUseCase := biz.NewTenantUserCase(tenantRepo, connStrGenerator, sender)
	permissionServiceServer := api2.NewPermissionGrpcClient(grpcConn)
	permissionChecker := api2.NewRemotePermissionChecker(permissionServiceServer)
	authzOption := server2.NewAuthorizationOption()
	subjectResolverImpl := authz.NewSubjectResolver(authzOption)
	defaultAuthorizationService := authz.NewDefaultAuthorizationService(permissionChecker, subjectResolverImpl, logger)
	factory := dal.NewBlobFactory(confData)
	tenantService := service.NewTenantService(tenantUseCase, defaultAuthorizationService, trustedContextValidator, factory, appConfig)
	httpServerRegister := service.NewHttpServerRegister(tenantService, factory, confData)
	httpServer := server2.NewHTTPServer(services, security, tokenizer, tenantStore, manager, webMultiTenancyOption, option, decodeRequestFunc, encodeResponseFunc, encodeErrorFunc, logger, trustedContextValidator, userTenantContributor, httpServerRegister)
	grpcServerRegister := service.NewGrpcServerRegister(tenantService)
	grpcServer := server2.NewGRPCServer(services, tokenizer, tenantStore, manager, webMultiTenancyOption, option, userTenantContributor, trustedContextValidator, grpcServerRegister, logger)
	migrate := data.NewMigrate(dataData)
	seeding := server2.NewSeeding(manager, migrate)
	seeder := server2.NewSeeder(seeding)
	tenantReadyEventHandler := biz.NewTenantReadyEventHandler(tenantUseCase)
	saasEventHandler := biz.NewRemoteEventHandler(logger, manager, tenantReadyEventHandler)
	handler := server2.NewEventHandler(saasEventHandler)
	receiver, cleanup5, err := dal.NewRemoteEventReceiver(confData, logger, handler, connName)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app := newApp(logger, httpServer, grpcServer, seeder, receiver)
	return app, func() {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

var (
	_wireEventBusValue           = eventbus.Default
	_wireConfigValue             = dal.UowCfg
	_wireDecodeRequestFuncValue  = server.ReqDecode
	_wireEncodeResponseFuncValue = server.ResEncoder
	_wireEncodeErrorFuncValue    = server.ErrEncoder
	_wireClientNameValue         = server2.ClientName
	_wireConnNameValue           = biz.ConnName
)
