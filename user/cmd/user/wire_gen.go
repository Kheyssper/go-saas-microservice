// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-saas/kit/pkg/api"
	"github.com/go-saas/kit/pkg/authn/jwt"
	"github.com/go-saas/kit/pkg/authz/authz"
	"github.com/go-saas/kit/pkg/authz/casbin"
	"github.com/go-saas/kit/pkg/conf"
	"github.com/go-saas/kit/pkg/dal"
	"github.com/go-saas/kit/pkg/gorm"
	"github.com/go-saas/kit/pkg/job"
	server2 "github.com/go-saas/kit/pkg/server"
	"github.com/go-saas/kit/pkg/uow"
	api2 "github.com/go-saas/kit/saas/api"
	api3 "github.com/go-saas/kit/user/api"
	"github.com/go-saas/kit/user/private/biz"
	conf2 "github.com/go-saas/kit/user/private/conf"
	"github.com/go-saas/kit/user/private/data"
	"github.com/go-saas/kit/user/private/server"
	"github.com/go-saas/kit/user/private/service"
	http2 "github.com/go-saas/kit/user/private/service/http"
	"github.com/go-saas/saas/http"
	"github.com/goxiaoy/go-eventbus"
)

import (
	_ "github.com/go-saas/kit/event/kafka"
	_ "github.com/go-saas/kit/event/pulsar"
	_ "github.com/go-saas/kit/pkg/registry/etcd"
	_ "github.com/go-saas/kit/saas/api"
	_ "github.com/go-saas/kit/sys/api"
)

// Injectors from wire.go:

// InitApp init kratos application.
func InitApp(services *conf.Services, security *conf.Security, userConf *conf2.UserConf, confData *conf.Data, logger log.Logger, webMultiTenancyOption *http.WebMultiTenancyOption, arg ...grpc.ClientOption) (*kratos.App, func(), error) {
	tokenizerConfig := jwt.NewTokenizerConfig(security)
	tokenizer := jwt.NewTokenizer(tokenizerConfig)
	dbCache, cleanup := gorm.NewDbCache(confData, logger)
	manager := uow.NewUowManager(dbCache)
	option := api.NewDefaultOption(logger)
	clientName := _wireClientNameValue
	discovery, err := api.NewDiscovery(services)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	inMemoryTokenManager := api.NewInMemoryTokenManager(tokenizer, logger)
	grpcConn, cleanup2 := api2.NewGrpcConn(clientName, services, discovery, option, inMemoryTokenManager, logger, arg...)
	tenantInternalServiceServer := api2.NewTenantInternalGrpcClient(grpcConn)
	tenantStore := api2.NewTenantStore(tenantInternalServiceServer)
	decodeRequestFunc := _wireDecodeRequestFuncValue
	encodeResponseFunc := _wireEncodeResponseFuncValue
	encodeErrorFunc := _wireEncodeErrorFuncValue
	connStrResolver := dal.NewConnStrResolver(confData, tenantStore)
	dbProvider := gorm.NewDbProvider(dbCache, connStrResolver, confData)
	dataData, cleanup3, err := data.NewData(confData, dbProvider, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData)
	passwordHasher := biz.NewPasswordHasher()
	userValidator := biz.NewUserValidator()
	passwordValidator := biz.NewPasswordValidator(userConf)
	lookupNormalizer := biz.NewLookupNormalizer()
	userTokenRepo := data.NewUserTokenRepo(dataData)
	refreshTokenRepo := data.NewRefreshTokenRepo(dataData)
	userTenantRepo := data.NewUserTenantRepo(dataData)
	connName := _wireConnNameValue
	client, err := dal.NewRedis(confData, connName)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	storeInterface := dal.NewCacheStore(client)
	cache := dal.NewStringCacheManager(storeInterface)
	emailTokenProvider := biz.NewEmailTokenProvider(cache)
	phoneTokenProvider := biz.NewPhoneTokenProvider(cache)
	userManager := biz.NewUserManager(userConf, userRepo, passwordHasher, userValidator, passwordValidator, lookupNormalizer, userTokenRepo, refreshTokenRepo, userTenantRepo, emailTokenProvider, phoneTokenProvider, cache, logger)
	eventBus := _wireEventBusValue
	roleRepo := data.NewRoleRepo(dataData, eventBus)
	roleManager := biz.NewRoleManager(roleRepo, lookupNormalizer)
	enforcerProvider, err := data.NewEnforcerProvider(logger, dbProvider)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	permissionService := casbin.NewPermissionService(enforcerProvider)
	userRoleContrib := service.NewUserRoleContrib(userRepo)
	authzOption := server.NewAuthorizationOption(userRoleContrib)
	subjectResolverImpl := authz.NewSubjectResolver(authzOption)
	defaultAuthorizationService := authz.NewDefaultAuthorizationService(permissionService, subjectResolverImpl, logger)
	factory := dal.NewBlobFactory(confData)
	trustedContextValidator := api.NewClientTrustedContextValidator()
	userService := service.NewUserService(userManager, roleManager, defaultAuthorizationService, factory, trustedContextValidator, logger)
	userTenantContrib := api3.NewUserTenantContrib(userService)
	lazyClient := dal.NewEmailer(confData)
	emailSender := biz.NewEmailSender(lazyClient, confData)
	authService := service.NewAuthService(userManager, roleManager, tokenizer, tokenizerConfig, passwordValidator, refreshTokenRepo, emailSender, security, defaultAuthorizationService, trustedContextValidator, logger)
	refreshTokenProvider := api3.NewRefreshProvider(authService, logger)
	tenantServiceServer := api2.NewTenantGrpcClient(grpcConn)
	userSettingRepo := data.NewUserSettingRepo(dataData, eventBus)
	userAddressRepo := data.NewUserAddrRepo(dataData, eventBus)
	accountService := service.NewAccountService(userManager, factory, tenantServiceServer, userSettingRepo, userAddressRepo, lookupNormalizer)
	roleService := service.NewRoleServiceService(roleManager, defaultAuthorizationService, permissionService)
	servicePermissionService := service.NewPermissionService(defaultAuthorizationService, permissionService, subjectResolverImpl, trustedContextValidator)
	signInManager := biz.NewSignInManager(userManager, security)
	apiClient := service.NewHydra(security)
	auth := http2.NewAuth(decodeRequestFunc, userManager, logger, signInManager, apiClient)
	httpServerRegister := service.NewHttpServerRegister(userService, encodeResponseFunc, encodeErrorFunc, accountService, authService, roleService, servicePermissionService, auth, confData, defaultAuthorizationService, factory)
	httpServer := server.NewHTTPServer(services, security, tokenizer, manager, webMultiTenancyOption, option, tenantStore, decodeRequestFunc, encodeResponseFunc, encodeErrorFunc, logger, userTenantContrib, trustedContextValidator, refreshTokenProvider, httpServerRegister)
	grpcServerRegister := service.NewGrpcServerRegister(userService, accountService, authService, roleService, servicePermissionService)
	grpcServer := server.NewGRPCServer(services, tokenizer, tenantStore, manager, webMultiTenancyOption, option, logger, trustedContextValidator, userTenantContrib, grpcServerRegister)
	redisConnOpt := job.NewAsynqClientOpt(client)
	migrate := data.NewMigrate(dataData)
	roleSeed := biz.NewRoleSeed(roleManager, permissionService)
	userSeed := biz.NewUserSeed(userManager, roleManager)
	permissionSeeder := biz.NewPermissionSeeder(permissionService, permissionService, roleManager)
	seeding := server.NewSeeding(manager, migrate, roleSeed, userSeed, permissionSeeder)
	seeder := server.NewSeeder(tenantStore, seeding)
	producer, cleanup4, err := dal.NewEventSender(confData, connName)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	userMigrationTaskHandler := biz.NewUserMigrationTaskHandler(seeder, producer)
	jobServer := server.NewJobServer(redisConnOpt, logger, userMigrationTaskHandler)
	asynqClient, cleanup5 := job.NewAsynqClient(redisConnOpt)
	tenantSeedEventHandler := biz.NewTenantSeedEventHandler(asynqClient)
	consumerFactoryServer := server.NewEventServer(confData, connName, logger, manager, tenantSeedEventHandler)
	registrar, err := server2.NewRegistrar(services)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app := newApp(userConf, logger, httpServer, grpcServer, jobServer, consumerFactoryServer, seeder, registrar)
	return app, func() {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

var (
	_wireClientNameValue         = server.ClientName
	_wireDecodeRequestFuncValue  = server2.ReqDecode
	_wireEncodeResponseFuncValue = server2.ResEncoder
	_wireEncodeErrorFuncValue    = server2.ErrEncoder
	_wireConnNameValue           = biz.ConnName
	_wireEventBusValue           = eventbus.Default
)
