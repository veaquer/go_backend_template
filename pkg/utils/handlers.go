package utils

import (
	"backend_template/pkg/errors/apperror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadBody[T any](c *gin.Context) (*T, bool) {
	var input T
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(apperror.New("Invalid input", http.StatusBadRequest))
		return nil, false
	}
	return &input, true
}

func ReadAndValidate[T any](c *gin.Context, validate func(T) error) (*T, bool) {
	var input T
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(apperror.New("Invalid input", http.StatusBadRequest))
		return nil, false
	}
	if err := validate(input); err != nil {
		c.Error(apperror.New(err.Error(), http.StatusBadRequest))
		return nil, false
	}
	return &input, true
}

func SetCookie(c *gin.Context, name, value string, maxAgeSeconds int, domain, path string, secure, httpOnly bool) {
	c.SetCookie(
		name,
		value,
		maxAgeSeconds,
		path,
		domain,
		secure,
		httpOnly,
	)
}

func SetRefreshCookie(c *gin.Context, token string) {
	SetCookie(c, "refresh_token", token, 3600*24*7, "localhost", "/", false, true)
}

func DeleteRefreshCookie(c *gin.Context) {
	DeleteCookie(c, "refresh_token", "localhost", "/")
}

func DeleteCookie(c *gin.Context, name, domain, path string) {
	c.SetCookie(
		name,
		"",
		-1,
		path,
		domain,
		false, // secure
		true,  // httpOnly
	)
}
