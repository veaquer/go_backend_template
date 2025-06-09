package token

import (
	"backend_template/internal/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager struct {
	accessSecret  string
	refreshSecret string
	accessExp     time.Duration
	refreshExp    time.Duration
}

type ExtractedToken struct {
	UserID uint
	Exp    int64
}

func NewTokenManager(cfg *config.Config) *TokenManager {
	return &TokenManager{
		accessSecret:  cfg.JWT.AccessSecret,
		refreshSecret: cfg.JWT.RefreshSecret,
		accessExp:     cfg.JWT.AccessExpiration,
		refreshExp:    cfg.JWT.RefreshExpiration,
	}
}

func (tk *TokenManager) GenerateAccessToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(tk.accessExp).Unix(),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(tk.accessSecret))
}

func (tk *TokenManager) GenereateRefreshToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(tk.refreshExp).Unix(),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(tk.refreshSecret))
}

func (tk *TokenManager) VerifyAccessToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(tk.accessSecret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return token, nil
}

func (tk *TokenManager) VerifyRefreshToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(tk.refreshSecret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return token, nil
}

func (tk *TokenManager) ExtractToken(token *jwt.Token) (uint, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token")
	}

	id, ok := claims["sub"]
	if !ok {
		return 0, errors.New("Missing subject in token claims")
	}

	idFloat, ok := id.(float64)
	if !ok {
		return 0, errors.New("Invalid subject type in token claims")
	}

	return uint(idFloat), nil
}
