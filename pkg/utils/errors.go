package utils

import (
	"github.com/veaquer/go_backend_template/pkg/errors/apperror"

	"github.com/gin-gonic/gin"
)

func AbortWithError(c *gin.Context, err *apperror.AppError) {
	c.AbortWithError(err.Code, err)
}

func AbortWithErrorNew(c *gin.Context, status int, msg string) {
	c.AbortWithError(status, apperror.New(msg, status))
}

func AbortWithErrorWrap(c *gin.Context, status int, msg string, err error) {
	c.AbortWithError(status, apperror.Wrap(msg, status, err))
}
