//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"nunu-layout-basic/internal/handler"
	"nunu-layout-basic/internal/repository"
	"nunu-layout-basic/internal/server"
	"nunu-layout-basic/internal/service"
	"nunu-layout-basic/pkg/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var RepositorySet = wire.NewSet(
	repository.NewDb,
	repository.NewRepository,
	repository.NewUserRepository,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

func newApp(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
	))
}
