package service

import (
	"order-service/internal/models/base"
	"order-service/internal/models/dto"
	"order-service/internal/models/entity"
	"order-service/internal/pkg/logging"
	"order-service/internal/repository"
)

type OrderService interface {
	CreateOrderWithLines(orderDTO dto.OrderDTO) (entity.TOrderModel, error)
	GetOrderByID(id uint) (entity.TOrderModel, error)
	GetAllOrders() ([]entity.TOrderModel, error)
}

type orderService struct {
	orderRepo     repository.OrderRepository
	orderLineRepo repository.OrderLineRepository
}

func NewOrderService(orderRepo repository.OrderRepository, orderLineRepo repository.OrderLineRepository) OrderService {
	return &orderService{orderRepo: orderRepo, orderLineRepo: orderLineRepo}
}

func (s *orderService) CreateOrderWithLines(orderDTO dto.OrderDTO) (entity.TOrderModel, error) {
	// Membuat Order
	order := entity.TOrderModel{
		Entity: base.Entity{
			Createdby: orderDTO.Createdby,
		},
		TOrderEntity: entity.TOrderEntity{
			TableID: orderDTO.TableID,
			Total:   orderDTO.Total,
		},
	}
	createdOrder, err := s.orderRepo.Create(order)
	if err != nil {
		logging.Log.Error("Error creating order:", err)
		return createdOrder, err
	}

	// Membuat Order Lines
	for _, lineDTO := range orderDTO.OrderLines {
		orderLine := entity.TOrderLinesModel{
			Entity: base.Entity{
				Createdby: orderDTO.Createdby,
			},
			TOrderLinesEntity: entity.TOrderLinesEntity{
				OrderID:  createdOrder.ID,
				MenuID:   lineDTO.MenuID,
				Quantity: lineDTO.Quantity,
				Price:    lineDTO.Price,
			},
		}
		subtotal := float64(orderLine.Quantity) * orderLine.Price
		orderLine.SubTotal = subtotal

		_, err := s.orderLineRepo.Create(orderLine)
		if err != nil {
			logging.Log.Error("Error creating order line:", err)
			return createdOrder, err
		}
	}

	return createdOrder, nil
}

func (s *orderService) GetOrderByID(id uint) (entity.TOrderModel, error) {
	order, err := s.orderRepo.FindByID(id)
	if err != nil {
		logging.Log.Error("Error finding order by ID:", err)
		return order, err
	}

	return order, nil
}

func (s *orderService) GetAllOrders() ([]entity.TOrderModel, error) {
	orders, err := s.orderRepo.FindAll()
	if err != nil {
		logging.Log.Error("Error finding all orders:", err)
		return orders, err
	}

	return orders, nil
}
