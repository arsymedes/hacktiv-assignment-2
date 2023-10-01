package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemCode    string `gorm:"not null;type:varchar(191)" json:"itemcode"`
	Description string `gorm:"not null;type:varchar(191)" json:"description"`
	Quantity    int    `gorm:"not null;type:int" json:"quantity"`
	OrderID     uint
}
