package domain

import (
	"postapp/pkg/model"
	"postapp/pkg/paginator"
	"time"
)

type Transaction struct {
	model.Base
	MerchantID uint64 `json:"merchant_id" gorm:"column:merchant_id"`
	OutletID   uint64 `json:"outlet_id" gorm:"column:outlet_id"`
	BillTotal  uint64 `json:"bill_total" gorm:"column:bill_total"`
}

type TransactionMerchant struct {
	TransactionDate string  `json:"transaction_date" gorm:"column:transaction_date"`
	Summary         float64 `json:"summary" gorm:"column:summary"`
}

type TransactionMerchantFilter struct {
	Pagination *paginator.Pagination `json:"pagination"`
	MerchantID uint64                `json:"merchant_id"`
	StartDate  *time.Time            `json:"start_date"`
	EndDate    *time.Time            `json:"end_date"`
}

type TransactionOutlet struct {
	TransactionDate string  `json:"transaction_date" gorm:"column:transaction_date"`
	Summary         float64 `json:"summary" gorm:"column:summary"`
}

type TransactionOutletFilter struct {
	Pagination *paginator.Pagination `json:"pagination"`
	OutletID   uint64                `json:"outlet_id"`
	StartDate  *time.Time            `json:"start_date"`
	EndDate    *time.Time            `json:"end_date"`
}
