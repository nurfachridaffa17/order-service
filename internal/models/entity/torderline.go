package entity

import (
	abstraction "order-service/internal/models/base"
)

type TOrderLinesEntity struct {
	ID       uint    `gorm:"primaryKey"`
	OrderID  uint    `gorm:"not null"`
	MenuID   uint    `gorm:"not null"`
	Quantity uint    `gorm:"not null"`
	Price    float64 `gorm:"not null"`
	SubTotal float64 `gorm:"not null"`
}

type TOrderLinesModel struct {
	TOrderLinesEntity
	abstraction.Entity
}

func (ol *TOrderLinesModel) TableName() string {
	return "t_order_lines"
}
