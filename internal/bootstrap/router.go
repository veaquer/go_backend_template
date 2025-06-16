package bootstrap

import (
	"github.com/veaquer/go_backend_template/internal/auth"
	"github.com/veaquer/go_backend_template/internal/auth/token"
	"github.com/veaquer/go_backend_template/internal/middleware"
	"github.com/veaquer/go_backend_template/internal/user/handler"

	"github.com/gin-gonic/gin"
)

func ProvideRouter(userHandler *handler.UserHandler, tk *token.TokenManager) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ErrorHandler())
	authMiddleware := auth.NewAuthMiddleware(tk)

	AuthGroup := r.Group("/auth")
	{
		AuthGroup.POST("/register", userHandler.Register)
		AuthGroup.POST("/login", userHandler.Login)
		AuthGroup.GET("/verify_email", userHandler.VerifyEmail)
		AuthGroup.POST("/refresh", authMiddleware.RequireRefreshToken(), userHandler.Refresh)
		AuthGroup.POST("/logout", userHandler.Logout)
	}

	user := r.Group("/user", authMiddleware.AuthRequired())
	{
		user.GET("/profile", userHandler.Profile)
	}

	return r
}
