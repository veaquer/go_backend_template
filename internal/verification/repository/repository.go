package repository

import (
	"backend_template/internal/verification/model"
	"context"
)

type VerificationRepository interface {
	CreateVerification(ctx context.Context, verification *model.EmailVerificationModel) error
	DeleteVerificationByID(ctx context.Context, ID uint) error
	GetVerificationByID(ctx context.Context, ID uint) (*model.EmailVerificationModel, error)
	GetVerificationByUserID(ctx context.Context, UserID uint) (*model.EmailVerificationModel, error)
	GetVerificationByToken(ctx context.Context, token string) (*model.EmailVerificationModel, error)
}
