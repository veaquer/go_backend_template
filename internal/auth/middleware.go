package auth

import (
	"backend_template/internal/auth/token"
	"backend_template/pkg/constants"
	"backend_template/pkg/errors/apperror"
	"backend_template/pkg/utils"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	tokens *token.TokenManager
}

func NewAuthMiddleware(tm *token.TokenManager) *AuthMiddleware {
	return &AuthMiddleware{tokens: tm}
}



func (m *AuthMiddleware) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.AbortWithError(c, apperror.ErrUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.AbortWithErrorNew(c, http.StatusUnauthorized, "Invalid auth header")
			return
		}

		tokenStr := parts[1]
		token, err := m.tokens.VerifyAccessToken(tokenStr)
		if err != nil {
			utils.AbortWithErrorWrap(c, http.StatusUnauthorized, "Invalid token", err)
			return
		}

		id, err := m.tokens.ExtractToken(token)
		if err != nil {
			utils.AbortWithErrorWrap(c, http.StatusUnauthorized, err.Error(), err)
			return
		}

		ctx := context.WithValue(c.Request.Context(), constants.UserIdCtxKey, id)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
