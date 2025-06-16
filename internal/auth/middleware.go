package auth

import (
	"github.com/veaquer/go_backend_template/internal/auth/token"
	"github.com/veaquer/go_backend_template/pkg/constants"
	"github.com/veaquer/go_backend_template/pkg/errors/apperror"
	"github.com/veaquer/go_backend_template/pkg/utils"
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
		c.Set(constants.UserIdCtxKey, id)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func (m *AuthMiddleware) RequireRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("refresh_token")
		if err != nil {
			utils.AbortWithErrorNew(c, http.StatusUnauthorized, "No refresh token")
			return
		}

		tok, err := m.tokens.VerifyRefreshToken(refreshToken)
		if err != nil {
			utils.AbortWithErrorNew(c, http.StatusUnauthorized, "Invalid or expired refresh token")
			return
		}

		id, err := m.tokens.ExtractToken(tok)
		if err != nil {
			utils.AbortWithErrorNew(c, http.StatusUnauthorized, "Failed to extract token")
			return
		}

		ctx := context.WithValue(c.Request.Context(), constants.UserIdCtxKey, id)
		c.Set(constants.UserIdCtxKey, id)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
