package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerName string    `gorm:"not null;type:varchar(191)" json:"customer_name"`
	Items        []Item    `json:"items"`
	OrderedAt    time.Time `json:"ordered_at"`
}

type NewOrderModel struct {
	gorm.Model
	CustomerName string    `gorm:"not null;type:varchar(191)" json:"customer_name" binding:"required"`
	Items        []Item    `json:"items" binding:"required"`
	OrderedAt    time.Time `json:"ordered_at" binding:"required"`
}
