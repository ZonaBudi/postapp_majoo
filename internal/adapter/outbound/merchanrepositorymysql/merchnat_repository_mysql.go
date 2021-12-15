package merchanrepositorymysql

import (
	"context"
	"errors"
	"postapp/internal/domain"
	"postapp/internal/port"

	"gorm.io/gorm"
)

type merchantRepoMysql struct {
	mysql *gorm.DB
}

func NewMerchantRepMysql(mysql *gorm.DB) port.MerchantRepository {
	return &merchantRepoMysql{
		mysql: mysql,
	}
}

func (_instance *merchantRepoMysql) FindOneByUserID(ctx context.Context, merchantID uint64) (*domain.Merchant, error) {
	var merchant domain.Merchant
	err := _instance.mysql.Where("user_id = ?", merchantID).First(&merchant).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &merchant, nil
}
