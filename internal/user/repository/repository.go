package repository

import (
	"github.com/veaquer/go_backend_template/internal/user/model"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.UserModel) error
	UpdateUser(ctx context.Context, user *model.UserModel) error
	GetUserByEmail(ctx context.Context, email string) (*model.UserModel, error)
	GetUserByID(ctx context.Context, id uint) (*model.UserModel, error)
	GetUserByUsername(ctx context.Context, username string) (*model.UserModel, error)
}
