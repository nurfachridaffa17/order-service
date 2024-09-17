package dto

import (
	model "order-service/internal/models/entity"
	res "order-service/internal/pkg/response"
)

type OrderDTO struct {
	TableID    uint32         `json:"table_id"`
	Total      float64        `json:"total"`
	Createdby  int            `json:"createdby"`
	OrderLines []OrderLineDTO `json:"order_lines"`
}

type OrderLineDTO struct {
	MenuID   uint    `json:"menu_id"`
	Quantity uint    `json:"quantity"`
	Price    float64 `json:"price"`
	SubTotal float64 `json:"sub_total"`
}

type (
	OrderResponse struct {
		model.TOrderModel
	}
	OrderResponseDoc struct {
		Body struct {
			Meta res.Meta      `json:"meta"`
			Data OrderResponse `json:"data"`
		} `json:"body"`
	}
)
