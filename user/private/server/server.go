package server

import (
	"github.com/google/wire"
	"github.com/goxiaoy/go-saas-kit/pkg/api"
	api2 "github.com/goxiaoy/go-saas-kit/user/api"
	"github.com/goxiaoy/go-saas-kit/user/private/biz"
	"github.com/goxiaoy/go-saas-kit/user/private/conf"
	"github.com/goxiaoy/go-saas-kit/user/private/data"
	seed2 "github.com/goxiaoy/go-saas-kit/user/private/seed"
	"github.com/goxiaoy/go-saas-kit/user/private/server/http"
	"github.com/goxiaoy/go-saas/seed"
	"github.com/goxiaoy/uow"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer, wire.Value(ClientName), seed2.NewFake, NewSeeder, http.NewAuth)

var ClientName api.ClientName = api2.ServiceName

func NewSeeder(c *conf.UserConf,
	uow uow.Manager,
	migrate *data.Migrate,
	roleSeed *biz.RoleSeed,
	userSeed *biz.UserSeed,
	fake *seed2.Fake,
	p *biz.PermissionSeeder) seed.Seeder {
	var opt = seed.NewSeedOption(migrate, seed.NewUowContributor(uow, seed.Chain(roleSeed, userSeed, fake, p)))
	// seed host
	opt.TenantIds = []string{""}

	return seed.NewDefaultSeeder(opt, map[string]interface{}{
		biz.AdminUsernameKey: c.Admin.GetUsername(),
		biz.AdminPasswordKey: c.Admin.GetPassword(),
		seed2.FakeSeedKey:    true,
	})
}
