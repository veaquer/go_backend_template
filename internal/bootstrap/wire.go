//go:build wireinject
// +build wireinject

package bootstrap

import (
	"backend_template/internal/auth/token"
	"backend_template/internal/config"
	"backend_template/internal/db"
	"backend_template/internal/logger"
	"backend_template/internal/user/repository"
	"backend_template/internal/user/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type App struct {
	Router *gin.Engine
}

func NewApp() (*App, error) {

	wire.Build(
		logger.New,
		config.Load,
		db.ProvideDB,
		repository.NewUserRepository,
		token.NewTokenManager,
		service.NewUserService,
		ProvideRouter,
		wire.Struct(new(App), "Router"),
	)

	return nil, nil
}
