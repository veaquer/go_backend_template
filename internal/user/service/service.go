package service

import (
	"github.com/veaquer/go_backend_template/internal/auth/token"
	"github.com/veaquer/go_backend_template/internal/cache"
	"github.com/veaquer/go_backend_template/internal/user/dto"
	"github.com/veaquer/go_backend_template/internal/user/model"
	"github.com/veaquer/go_backend_template/internal/user/repository"
	"github.com/veaquer/go_backend_template/internal/verification/service"
	"github.com/veaquer/go_backend_template/pkg/errors/apperror"
	"github.com/veaquer/go_backend_template/pkg/hash"
	"context"
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct {
	repo                repository.UserRepository
	logger              *zap.Logger
	verificationService *service.VerificationService
	tk                  *token.TokenManager
	redis               *cache.RedisCache
}

func NewUserService(r repository.UserRepository, l *zap.Logger, verificationService *service.VerificationService, tk *token.TokenManager, redis *cache.RedisCache) *UserService {
	return &UserService{repo: r, logger: l, verificationService: verificationService, tk: tk, redis: redis}
}

func (s *UserService) Register(ctx context.Context, input dto.RegisterUserDto) error {
	// Check if username or email taken
	existUsername, err := s.repo.GetUserByUsername(ctx, input.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.Wrap("Failed to check username existence", 500, err)
	}

	if existUsername.ID != 0 {
		return apperror.NewConflict("username", "This username is already taken")
	}

	existEmail, err := s.repo.GetUserByEmail(ctx, input.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.Wrap("Failed to check email existence", 500, err)
	}
	if existEmail.ID != 0 {
		return apperror.NewConflict("email", "This email is already taken")
	}

	hashedPassword, err := hash.HashPassword(input.Password)
	if err != nil {
		return apperror.Wrap("Failed to hash password", 500, err)
	}

	user := model.UserModel{Name: input.Name, Username: input.Username,
		Email: input.Email, Password: *hashedPassword}

	err = s.repo.CreateUser(ctx, &user)
	if err != nil {
		return apperror.Wrap("Failed to create user", 500, err)
	}

	if err := s.verificationService.SendVerification(ctx, user.ID, user.Email, "register"); err != nil {
		return apperror.Wrap("Failed to send email verification", 500, err)
	}

	return nil
}

func (s *UserService) Login(ctx context.Context, input dto.LoginUserDto) (*model.Tokens, error) {
	user, err := s.repo.GetUserByUsername(ctx, input.Username)
	if err != nil || user.ID == 0 {
		return nil, apperror.New("Invalid username", http.StatusUnauthorized)
	}

	if err := hash.ComparePassword(user.Password, input.Password); err != nil {
		return nil, apperror.New("Invalid password", http.StatusUnauthorized)
	}

	if !user.IsEmailVerified {
		return nil, apperror.New("Email is not verified", http.StatusUnauthorized)
	}

	accessToken, err := s.tk.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, apperror.NewInternal("Failed to generate access token")
	}

	refreshToken, err := s.tk.GenereateRefreshToken(user.ID)
	if err != nil {
		return nil, apperror.NewInternal("Failed to generate refresh token")
	}

	tokens := model.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}
	return &tokens, nil
}

func (s *UserService) Refresh(ctx context.Context, id uint) (*model.Tokens, error) {
	accessToken, err := s.tk.GenerateAccessToken(id)
	if err != nil {
		return nil, apperror.NewInternal("Failed to generate access token")
	}

	refreshToken, err := s.tk.GenereateRefreshToken(id)
	if err != nil {
		return nil, apperror.NewInternal("Failed to generate refresh token")
	}

	tokens := model.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}
	return &tokens, nil
}

func (s *UserService) GetUserById(ctx context.Context, id uint) (*model.UserModel, error) {
	// Redis
	user, err := s.redis.GetUserProfile(ctx, id)
	if err != nil {
		return nil, apperror.NewInternal("Redis failure")
	}

	if user != nil {
		return user, nil
	}

	user, err = s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, apperror.Wrap("Unauthorized", http.StatusUnauthorized, err)
	}

	_ = s.redis.SetUserProfile(ctx, user, 10*time.Minute)

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, updatedUser *model.UserModel) error {
	return s.repo.UpdateUser(ctx, updatedUser)
}
