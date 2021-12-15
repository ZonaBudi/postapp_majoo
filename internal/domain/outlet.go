package domain

import "postapp/pkg/model"

type Outlet struct {
	model.Base
	MerchantID uint64 `json:"merchant_id" gorm:"column:merchant_id"`
	OutletName string `json:"outlet_name" gorm:"column:outlet_name"`
}

func (u *Outlet) IsEmpty() bool {
	return u == nil
}

func (u *Outlet) MyOutlet(MerchantID uint64) bool {
	return u.MerchantID == MerchantID
}

type PublicOutlet struct {
	OutletName string `json:"outlet_name" gorm:"column:outlet_name"`
}
