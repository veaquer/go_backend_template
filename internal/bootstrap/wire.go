//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/veaquer/go_backend_template/internal/auth/token"
	"github.com/veaquer/go_backend_template/internal/config"
	"github.com/veaquer/go_backend_template/internal/db"
	"github.com/veaquer/go_backend_template/internal/logger"
	"github.com/veaquer/go_backend_template/internal/user/handler"
	userRepo "github.com/veaquer/go_backend_template/internal/user/repository"
	userService "github.com/veaquer/go_backend_template/internal/user/service"
	"github.com/veaquer/go_backend_template/internal/verification/repository"
	"github.com/veaquer/go_backend_template/internal/verification/service"
	"github.com/veaquer/go_backend_template/pkg/email"

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

var providers = wire.NewSet(
		ProvideRedis,
		ProvideRouter,
		db.ProvideDB,
		)

var repositories = wire.NewSet(
		userRepo.NewUserRepository,
		repository.NewVerificationRepository,
		)

var services = wire.NewSet(
		userService.NewUserService,
		service.NewVerificationService,
		)

var handlers = wire.NewSet(
		handler.NewUserHandler,
		)

func NewApp() (*App, error) {

	wire.Build(
		logger.New,
		config.Load,
		token.NewTokenManager,
		email.NewGoMailSender,
		providers,
		repositories,
		services,
		handlers,
		wire.Struct(new(App), "Router", "Logger", "DB"),
	)

	return nil, nil
}
