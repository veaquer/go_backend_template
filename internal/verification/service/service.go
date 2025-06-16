package service

import (
	"github.com/veaquer/go_backend_template/internal/config"
	UserModel "github.com/veaquer/go_backend_template/internal/user/model"
	"github.com/veaquer/go_backend_template/internal/verification/model"
	"github.com/veaquer/go_backend_template/internal/verification/repository"
	"github.com/veaquer/go_backend_template/pkg/constants"
	"github.com/veaquer/go_backend_template/pkg/email"
	"github.com/veaquer/go_backend_template/pkg/errors/apperror"

	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type VerificationService struct {
	repo   repository.VerificationRepository
	sender *email.GoMailSender
	cfg    *config.Config
}

func NewVerificationService(repo repository.VerificationRepository, sender *email.GoMailSender, cfg *config.Config) *VerificationService {
	return &VerificationService{repo: repo, sender: sender, cfg: cfg}
}

func (s *VerificationService) SendVerification(ctx context.Context, UserID uint, Email, Purpose string) error {

	uuid := uuid.New().String()
	newVerification := model.EmailVerificationModel{
		UserID: UserID, Email: Email, Purpose: Purpose,
		Token:     uuid,
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(constants.VerificationTokenExpiry)),
		CreatedAt: time.Now(),
	}

	var URL string
	if s.cfg.Env == "development" && len(s.cfg.BackendUrl) == 0 {
		URL = "http://localhost:" + s.cfg.Port
	} else {
		URL = s.cfg.BackendUrl
	}

	body := "<a href='" + URL + constants.FullVerificationPath + "?token=" + uuid + "'>Verify</a>"
	mail := email.Email{
		To:      Email,
		Subject: "Email Verification",
		Body:    body,
		IsHTML:  true,
	}
	if err := s.sender.SendEmail(mail); err != nil {
		return apperror.New("Failed to send verification email: %w", http.StatusBadRequest)
	}

	err := s.repo.CreateVerification(ctx, &newVerification)
	if err != nil {
		return apperror.New("Failed to create verification record: %w", http.StatusBadRequest)
	}

	return nil
}

func (s *VerificationService) Validate(ctx context.Context, token string) (*UserModel.UserModel, error) {
	verification, err := s.repo.GetVerificationByToken(ctx, token)
	if err != nil {
		return nil, apperror.New("Invalid or expired token", http.StatusBadRequest)
	}

	if verification.ExpiresAt.Before(time.Now()) {
		s.repo.DeleteVerificationByID(ctx, verification.ID)
		return nil, apperror.New("Expired token", http.StatusBadRequest)
	}

	return &UserModel.UserModel{ID: verification.UserID, IsEmailVerified: true, Email: verification.Email}, nil
}
