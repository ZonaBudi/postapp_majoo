package domain

import (
	"crypto/md5"
	"encoding/hex"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Login struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (lr LoginRequest) LoginValidation() error {
	return validation.ValidateStruct(&lr,
		validation.Field(&lr.Username, validation.Required),
		validation.Field(&lr.Password, validation.Required),
	)
}

func (l *LoginRequest) HashedPassword() string {
	hasher := md5.New()
	hasher.Write([]byte(l.Password))
	incomingPassword := hex.EncodeToString(hasher.Sum(nil))
	return incomingPassword
}
