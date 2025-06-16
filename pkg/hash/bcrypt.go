package hash

import (
	"github.com/veaquer/go_backend_template/pkg/constants"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (*string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), constants.Cost)
	if err != nil {
		return nil, err
	}

	str := string(hash)
	return &str, nil
}

func ComparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
