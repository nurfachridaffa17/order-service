package response

import "order-service/internal/models/base"

type Meta struct {
	Success bool                 `json:"success" default:"true"`
	Message string               `json:"message" default:"true"`
	Info    *base.PaginationInfo `json:"info"`
}
