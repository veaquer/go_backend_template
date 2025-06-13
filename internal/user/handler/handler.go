package handler

import (
	"backend_template/internal/user/dto"
	UserService "backend_template/internal/user/service"
	"backend_template/internal/verification/service"
	"backend_template/pkg/constants"
	"backend_template/pkg/errors/apperror"
	"backend_template/pkg/utils"
	"backend_template/pkg/validator"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService         *UserService.UserService
	verificationService *service.VerificationService
}

func NewUserHandler(s *UserService.UserService, v *service.VerificationService) *UserHandler {
	return &UserHandler{userService: s, verificationService: v}
}

func (h *UserHandler) Register(c *gin.Context) {
	input, ok := utils.ReadAndValidate(c, validator.ValidateStruct[dto.RegisterUserDto])
	if !ok {
		return
	}

	if err := h.userService.Register(c, *input); err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, gin.H{
		"message": "Successfully registered!\nWe've sent you verification link to email",
	})

}

func (h *UserHandler) Login(c *gin.Context) {
	input, ok := utils.ReadAndValidate(c, validator.ValidateStruct[dto.LoginUserDto])
	if !ok {
		return
	}

	tokens, err := h.userService.Login(c, *input)
	if err != nil {
		c.Error(err)
		return
	}

	utils.SetRefreshCookie(c, tokens.RefreshToken)
	c.JSON(200, gin.H{
		"accessToken": tokens.AccessToken,
	})
}

func (h *UserHandler) Refresh(c *gin.Context) {
	userID := c.GetUint(constants.UserIdCtxKey)

	tokens, err := h.userService.Refresh(c, userID)
	if err != nil {
		c.Error(err)
		return
	}

	utils.SetRefreshCookie(c, tokens.RefreshToken)
	c.JSON(200, gin.H{
		"accessToken": tokens.AccessToken,
	})
}

func (h *UserHandler) Logout(c *gin.Context) {
	utils.DeleteRefreshCookie(c)
	c.JSON(200, gin.H{
		"message": "Successfully logged out!",
	})
}

func (h *UserHandler) Profile(c *gin.Context) {
	userID := c.GetUint(constants.UserIdCtxKey)
	user, err := h.userService.GetUserById(c, userID)
	if err != nil {
		c.Error(err)
	}

	c.JSON(200, user)
}

func (h *UserHandler) VerifyEmail(c *gin.Context) {
	user, err := h.verificationService.Validate(c, c.Query("token"))
	if err != nil {
		c.Error(err)
		return
	}

	err = h.userService.UpdateUser(c, user)
	if err != nil {
		c.Error(apperror.Wrap("Failed to update user", 500, err))
	}

	c.String(200, "Successfully verified email!")
}
