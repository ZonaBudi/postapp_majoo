package outletrepositorymysql

import (
	"context"
	"errors"
	"postapp/internal/domain"
	"postapp/internal/port"

	"gorm.io/gorm"
)

type outletRepoMysql struct {
	mysql *gorm.DB
}

func NewOutletRepMysql(mysql *gorm.DB) port.OutletRepository {
	return &outletRepoMysql{
		mysql: mysql,
	}
}

func (_instance *outletRepoMysql) FindOneByOutletID(ctx context.Context, outletID uint64) (*domain.Outlet, error) {
	var outlet domain.Outlet
	err := _instance.mysql.Where("id = ?", outletID).First(&outlet).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &outlet, nil
}
