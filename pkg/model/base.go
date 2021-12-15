package model

import "time"

type Base struct {
	ID        *uint64   `json:"id" gorm:"column:id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	CreatedBy string    `json:"created_by" gorm:"column:created_by"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy string    `json:"updated_by" gorm:"column:updated_by"`
}

var (
	DateFormatFilter = "2006-01-02"
)
