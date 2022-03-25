package helpers

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Password struct{}

func NewPassword() *Password {
	return &Password{}
}

func (p Password) Encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.Wrap(err, "Error en generación de password")
	}
	return string(bytes), nil
}

func (p Password) Validate(rqPass string, compPass string) (bool, error) {
	passRequestBytes := []byte(rqPass)
	compPassBytes := []byte(compPass)

	err := bcrypt.CompareHashAndPassword(compPassBytes, passRequestBytes)
	if err != nil {
		return false, err
	}
	return true, nil
}
