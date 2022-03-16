package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/goxiaoy/go-saas-kit/pkg/authz/casbin"
	"github.com/goxiaoy/go-saas-kit/pkg/blob"
	gorm2 "github.com/goxiaoy/go-saas-kit/pkg/gorm"
	uow2 "github.com/goxiaoy/go-saas-kit/pkg/uow"
	"github.com/goxiaoy/go-saas-kit/user/private/biz"
	"github.com/goxiaoy/go-saas-kit/user/private/conf"
	"github.com/goxiaoy/go-saas/common"
	"github.com/goxiaoy/go-saas/data"
	"github.com/goxiaoy/go-saas/gorm"
	g "gorm.io/gorm"

	_ "github.com/goxiaoy/go-saas-kit/pkg/blob/memory"
	_ "github.com/goxiaoy/go-saas-kit/pkg/blob/os"
	_ "github.com/goxiaoy/go-saas-kit/pkg/blob/s3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	gorm.NewDbOpener,
	uow2.NewUowManager,
	NewBlobFactory,
	NewProvider,
	NewEnforcerProvider,
	NewUserRepo,
	NewRefreshTokenRepo,
	NewRoleRepo,
	NewUserTenantRepo,
	NewMigrate,
	NewUserTokenRepo,
	NewUserSettingRepo)

const ConnName = "user"

// Data .
type Data struct {
	DbProvider gorm.DbProvider
}

func GetDb(ctx context.Context, provider gorm.DbProvider) *g.DB {
	db := provider.Get(ctx, ConnName)
	if err := db.SetupJoinTable(&biz.User{}, "Roles", &biz.UserRole{}); err != nil {
		panic(err)
	}
	gorm2.RegisterCallbacks(db)
	return db
}

// NewData .
func NewData(c *conf.Data, dbProvider gorm.DbProvider, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		logger.Log(log.LevelInfo, "closing the data resources")
	}
	return &Data{
		DbProvider: dbProvider,
	}, cleanup, nil
}

func NewProvider(c *conf.Data, cfg *gorm.Config, opener gorm.DbOpener, ts common.TenantStore, logger log.Logger) gorm.DbProvider {
	conn := make(data.ConnStrings, 1)
	for k, v := range c.Endpoints.Databases {
		conn[k] = v.Source
	}
	mr := common.NewMultiTenancyConnStrResolver(func() common.TenantStore {
		return ts
	}, data.NewConnStrOption(conn))
	r := gorm.NewDefaultDbProvider(mr, cfg, opener)
	return r
}

func NewEnforcerProvider(dbProvider gorm.DbProvider) *casbin.EnforcerProvider {
	return casbin.NewEnforcerProvider(dbProvider, ConnName)
}

func NewBlobFactory(c *conf.Data) blob.Factory {
	return blob.NewFactory(c.Blobs)
}
