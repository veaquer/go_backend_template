package bootstrap

import (
	"backend_template/internal/auth/token"
	"backend_template/internal/middleware"
	"backend_template/internal/user/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ProvideRouter(userSvc *service.UserService, tokenMgr *token.TokenManager) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ErrorHandler())
	fmt.Println("Set up router")
	return r
}
