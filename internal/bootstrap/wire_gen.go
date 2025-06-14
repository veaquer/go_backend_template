// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bootstrap

import (
	"backend_template/internal/auth/token"
	"backend_template/internal/config"
	"backend_template/internal/db"
	"backend_template/internal/logger"
	"backend_template/internal/user/handler"
	"backend_template/internal/user/repository"
	service2 "backend_template/internal/user/service"
	repository2 "backend_template/internal/verification/repository"
	"backend_template/internal/verification/service"
	"backend_template/pkg/email"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func NewApp() (*App, error) {
	zapLogger := logger.New()
	configConfig := config.Load(zapLogger)
	gormDB, err := db.ProvideDB(configConfig)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	verificationRepository := repository2.NewVerificationRepository(gormDB)
	goMailSender := email.NewGoMailSender(configConfig)
	verificationService := service.NewVerificationService(verificationRepository, goMailSender, configConfig)
	tokenManager := token.NewTokenManager(configConfig)
	userService := service2.NewUserService(userRepository, zapLogger, verificationService, tokenManager)
	redisCache := ProvideRedis(configConfig)
	userHandler := handler.NewUserHandler(userService, verificationService, redisCache)
	engine := ProvideRouter(userHandler, tokenManager)
	app := &App{
		Router: engine,
		Logger: zapLogger,
		DB:     gormDB,
	}
	return app, nil
}

// wire.go:

type App struct {
	Router *gin.Engine
	Logger *zap.Logger
	DB     *gorm.DB
}
