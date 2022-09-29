//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"

	"Test/internal/app/api"
	"Test/internal/app/dao"
	"Test/internal/app/router"
	"Test/internal/app/service"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGormDB,
		dao.RepoSet,
		InitAuth,
		InitGinEngine,
		service.ServiceSet,
		api.APISet,
		router.RouterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
