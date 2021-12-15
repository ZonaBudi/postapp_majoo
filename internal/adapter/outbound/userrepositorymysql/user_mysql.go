package userrepositorymysql

import (
	"context"
	"errors"
	"postapp/internal/domain"
	"postapp/internal/port"

	"gorm.io/gorm"
)

type userRepoMysql struct {
	mysql *gorm.DB
}

func NewUserRepMysql(mysql *gorm.DB) port.UserRepository {
	return &userRepoMysql{
		mysql: mysql,
	}
}

func (_instance *userRepoMysql) FindOneByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	err := _instance.mysql.Where("user_name = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
