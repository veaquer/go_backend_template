package repository

import (
	"github.com/veaquer/go_backend_template/internal/verification/model"
	"context"

	"gorm.io/gorm"
)

type verificationRepository struct {
	db *gorm.DB
}

func NewVerificationRepository(db *gorm.DB) VerificationRepository {
	return &verificationRepository{db: db}
}

func (r *verificationRepository) CreateVerification(ctx context.Context, verification *model.EmailVerificationModel) error {
	return r.db.WithContext(ctx).Save(verification).Error
}

func (r *verificationRepository) GetVerificationByID(ctx context.Context, ID uint) (*model.EmailVerificationModel, error) {
	var verification model.EmailVerificationModel
	err := r.db.WithContext(ctx).Where("id = ?", ID).First(&verification).Error

	return &verification, err
}

func (r *verificationRepository) GetVerificationByUserID(ctx context.Context, UserID uint) (*model.EmailVerificationModel, error) {
	var verification model.EmailVerificationModel
	err := r.db.WithContext(ctx).Where("user_id = ?", UserID).First(&verification).Error

	return &verification, err
}

func (r *verificationRepository) GetVerificationByToken(ctx context.Context, token string) (*model.EmailVerificationModel, error) {
	var verification model.EmailVerificationModel
	err := r.db.WithContext(ctx).Where("token = ?", token).First(&verification).Error
	return &verification, err
}

func (r *verificationRepository) DeleteVerificationByID(ctx context.Context, ID uint) error {
	return r.db.WithContext(ctx).Delete(&model.EmailVerificationModel{ID: ID}).Error
}
