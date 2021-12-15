package domain

import "postapp/pkg/model"

type Merchant struct {
	model.Base
	UserID       uint64 `json:"-"  gorm:"column:user_id"`
	MerchantName string `json:"merchant_name" gorm:"column:merchant_name"`
}

func (u *Merchant) IsEmpty() bool {
	return u == nil
}

type PublicMerchant struct {
	MerchantName string `json:"merchant_name" gorm:"column:merchant_name"`
}
