package entity

import (
	abstraction "order-service/internal/models/base"
)

type TOrderEntity struct {
	ID         uint               `gorm:"primaryKey"`
	TableID    uint32             `gorm:"not null"`
	Total      float64            `gorm:"not null"`
	OrderLines []TOrderLinesModel `gorm:"foreignKey:OrderID"`
}

type TOrderModel struct {
	TOrderEntity
	abstraction.Entity
}

func (o *TOrderModel) TableName() string {
	return "t_orders"
}
