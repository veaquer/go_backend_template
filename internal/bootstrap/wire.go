//go:build wireinject
// +build wireinject

package bootstrap

import (
	"backend_template/internal/auth/token"
	"backend_template/internal/config"
	"backend_template/internal/db"
	"backend_template/internal/logger"
	"backend_template/internal/user/handler"
	userRepo "backend_template/internal/user/repository"
	userService "backend_template/internal/user/service"
	"backend_template/internal/verification/repository"
	"backend_template/internal/verification/service"
	"backend_template/pkg/email"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	Router *gin.Engine
	Logger *zap.Logger
	DB     *gorm.DB
}

func NewApp() (*App, error) {

	wire.Build(
		logger.New,
		config.Load,
		db.ProvideDB,
		userRepo.NewUserRepository,
		token.NewTokenManager,
		email.NewGoMailSender,
		repository.NewVerificationRepository,
		service.NewVerificationService,
		ProvideRedis,
		userService.NewUserService,
		handler.NewUserHandler,
		ProvideRouter,
		wire.Struct(new(App), "Router", "Logger", "DB"),
	)

	return nil, nil
}
