package handler

import (
	"order-service/internal/controller"
	"order-service/internal/repository"
	"order-service/internal/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type handlerOrder struct {
	Controller controller.OrderController
}

func NewOrderHandler(db *gorm.DB) *handlerOrder {
	or := repository.NewOrderRepository(db)
	orl := repository.NewOrderLineRepository(db)
	cs := service.NewOrderService(or, orl)

	return &handlerOrder{
		Controller: controller.NewOrderController(cs),
	}
}

func (h *handlerOrder) Route(g *echo.Group) {
	g.GET("/orders", h.Controller.GetOrders)
}
