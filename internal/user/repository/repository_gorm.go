package repository

import (
	"backend_template/internal/user/model"
	"context"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.UserModel) error {
	return r.db.WithContext(ctx).Model(&model.UserModel{}).Where("id = ?", user.ID).Updates(user).Error
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.UserModel) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.UserModel, error) {
	var user model.UserModel
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) GetUserByID(ctx context.Context, id uint) (*model.UserModel, error) {
	var user model.UserModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*model.UserModel, error) {
	var user model.UserModel
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	return &user, err
}
