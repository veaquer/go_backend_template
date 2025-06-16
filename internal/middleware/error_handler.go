package middleware

import (
	"github.com/veaquer/go_backend_template/pkg/errors/apperror"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, ginErr := range c.Errors {
			if err, ok := ginErr.Err.(apperror.ErrorResponder); ok {
				c.JSON(err.GetCode(), gin.H{"error": err.GetMessage()})
				return
			}
			// fallback for unexpected errors
			c.JSON(500, gin.H{"error": "Internal Server Error"})
		}
	}
}
