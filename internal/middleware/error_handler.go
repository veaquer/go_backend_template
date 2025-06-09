package middleware

import (
	"backend_template/pkg/errors/apperror"
	"errors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, ginErr := range c.Errors {
			var appErr *apperror.AppError
			if errors.As(ginErr.Err, &appErr) {
				c.JSON(appErr.Code, gin.H{"error": appErr.Message})
			}
		}
	}
}
