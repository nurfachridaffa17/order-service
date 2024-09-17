package controller

import (
	"github.com/labstack/echo/v4"

	"order-service/internal/models/dto"
	"order-service/internal/service"

	res "order-service/internal/pkg/response"
)

type OrderController interface {
	// GetOrders
	GetOrders(c echo.Context) error
	// GetOrderID
	// GetOrderByID(c echo.Context) error
	// // CreateOrder
	// CreateOrder(c echo.Context) error
	// // UpdateOrder
	// UpdateOrder(c echo.Context) error
	// // DeleteOrder
	// DeleteOrder(c echo.Context) error
}

type orderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &orderController{orderService}
}

func (h *orderController) GetOrders(c echo.Context) error {
	var response []dto.OrderResponse
	result, err := h.orderService.GetAllOrders()
	if err != nil {
		return res.CustomErrorBuilder(500, err.Error(), "Failed to get orders").Send(c)
	}

	for _, v := range result {
		response = append(response, dto.OrderResponse{
			TOrderModel: v,
		})
	}

	return res.SuccessResponse(response).Send(c)
}
