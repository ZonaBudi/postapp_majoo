package domain

import (
	"errors"
	"postapp/pkg/model"

	"github.com/golang-jwt/jwt"
)

type User struct {
	model.Base
	Name     string `json:"name" gorm:"column:name"`
	UserName string `json:"user_name" gorm:"column:user_name"`
	Password string `json:"-" gorm:"column:password"`
}

func (u *User) IsEmpty() bool {
	return u == nil
}
func (u *User) ComparePassword(password string) bool {
	return u.Password == password
}

func (u *User) GenerateTokenAccess(secret string) (*string, error) {
	mySigningKey := []byte(secret)
	// Create the Claims
	if u.ID == nil {
		return nil, errors.New("user id is nil")
	}
	claims := &jwt.MapClaims{
		"uid": u.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return nil, err
	}
	return &ss, nil
}
