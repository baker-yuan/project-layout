//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"nunu-layout-advanced/internal/handler"
	"nunu-layout-advanced/internal/middleware"
	"nunu-layout-advanced/internal/repository"
	"nunu-layout-advanced/internal/server"
	"nunu-layout-advanced/internal/service"
	"nunu-layout-advanced/pkg/helper/sid"
	"nunu-layout-advanced/pkg/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var SidSet = wire.NewSet(sid.NewSid)

var JwtSet = wire.NewSet(middleware.NewJwt)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var RepositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewUserRepository,
)

func newApp(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
		SidSet,
		JwtSet,
	))
}
