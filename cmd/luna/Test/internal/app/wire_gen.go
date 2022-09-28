// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"Test/internal/app/api"
	"Test/internal/app/dao/Greet"
	"Test/internal/app/dao/util"
	"Test/internal/app/router"
	"Test/internal/app/service"
)

// Injectors from wire.go:

func BuildInjector() (*Injector, func(), error) {
	auther, cleanup, err := InitAuth()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup2, err := InitGormDB()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	trans := &util.Trans{
		DB: db,
	}
	greetRepo := &Greet.GreetRepo{
		DB: db,
	}
	greetSrv := &service.GreetSrv{
		TransRepo: trans,
		GreetRepo: greetRepo,
	}
	greetAPI := &api.GreetAPI{
		GreetSrv: greetSrv,
	}
	routerRouter := &router.Router{
		Auth:     auther,
		GreetAPI: greetAPI,
	}
	engine := InitGinEngine(routerRouter)
	injector := &Injector{
		Engine: engine,
		Auth:   auther,
	}
	return injector, func() {
		cleanup2()
		cleanup()
	}, nil
}