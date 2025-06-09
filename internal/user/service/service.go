package service

import (
	"backend_template/internal/user/repository"

	"go.uber.org/zap"
)

type UserService struct {
	repo   repository.UserRepository
	logger *zap.Logger
}

func NewUserService(r repository.UserRepository, l *zap.Logger) *UserService {
	return &UserService{repo: r, logger: l}
}
