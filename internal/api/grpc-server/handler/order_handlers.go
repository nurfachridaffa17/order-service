package handler

import (
	"context"
	"order-service/internal/models/dto"
	"order-service/internal/pkg/logging"
	"order-service/internal/service"
	order_service "order-service/proto/order-service/proto/order"
)

type OrderHandler struct {
	orderService service.OrderService
	order_service.UnimplementedOrderServiceServer
}

func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

// Implement the CreateOrder method for gRPC
func (h *OrderHandler) CreateOrder(ctx context.Context, req *order_service.OrderRequest) (*order_service.OrderResponse, error) {
	// Convert gRPC request to DTO
	orderDTO := dto.OrderDTO{
		TableID:    req.TableId,
		Total:      float64(req.Total),
		Createdby:  int(req.CreatedBy),
		OrderLines: []dto.OrderLineDTO{},
	}

	for _, line := range req.OrderLines {
		orderLine := dto.OrderLineDTO{
			MenuID:   uint(line.MenuId),
			Quantity: uint(line.Quantity),
			Price:    float64(line.Price),
		}
		orderDTO.OrderLines = append(orderDTO.OrderLines, orderLine)
	}

	// Call service layer to create the order with lines
	createdOrder, err := h.orderService.CreateOrderWithLines(orderDTO)
	if err != nil {
		logging.Log.Error("Error creating order:", err)
		return nil, err
	}

	// Convert the result back to gRPC response
	orderResponse := &order_service.OrderResponse{
		Id:         uint32(createdOrder.ID),
		TableId:    createdOrder.TableID,
		Total:      float32(createdOrder.Total),
		OrderLines: []*order_service.OrderLineResponse{},
	}

	for _, line := range createdOrder.OrderLines {
		orderLineResponse := &order_service.OrderLineResponse{
			MenuId:   uint32(line.MenuID),
			Quantity: uint32(line.Quantity),
			Price:    float32(line.Price),
			SubTotal: float32(line.SubTotal),
		}
		orderResponse.OrderLines = append(orderResponse.OrderLines, orderLineResponse)
	}

	return orderResponse, nil
}

// Implement the GetOrder method
func (h *OrderHandler) GetOrder(ctx context.Context, req *order_service.OrderIDRequest) (*order_service.OrderResponse, error) {
	order, err := h.orderService.GetOrderByID(uint(req.OrderId))
	if err != nil {
		logging.Log.Error("Error finding order:", err)
		return nil, err
	}

	// Convert the result back to gRPC response
	orderResponse := &order_service.OrderResponse{
		Id:         uint32(order.ID),
		TableId:    order.TableID,
		Total:      float32(order.Total),
		OrderLines: []*order_service.OrderLineResponse{},
	}

	for _, line := range order.OrderLines {
		orderLineResponse := &order_service.OrderLineResponse{
			MenuId:   uint32(line.MenuID),
			Quantity: uint32(line.Quantity),
			Price:    float32(line.Price),
			SubTotal: float32(line.SubTotal),
		}
		orderResponse.OrderLines = append(orderResponse.OrderLines, orderLineResponse)
	}

	return orderResponse, nil
}

// Implement the GetAllOrders method
func (h *OrderHandler) GetAllOrders(ctx context.Context, req *order_service.Empty) (*order_service.OrderListResponse, error) {
	orders, err := h.orderService.GetAllOrders()
	if err != nil {
		logging.Log.Error("Error finding all orders:", err)
		return nil, err
	}

	orderListResponse := &order_service.OrderListResponse{
		Orders: []*order_service.OrderResponse{},
	}

	for _, order := range orders {
		orderResponse := &order_service.OrderResponse{
			Id:         uint32(order.ID),
			TableId:    order.TableID,
			Total:      float32(order.Total),
			OrderLines: []*order_service.OrderLineResponse{},
		}

		for _, line := range order.OrderLines {
			orderLineResponse := &order_service.OrderLineResponse{
				MenuId:   uint32(line.MenuID),
				Quantity: uint32(line.Quantity),
				Price:    float32(line.Price),
				SubTotal: float32(line.SubTotal),
			}
			orderResponse.OrderLines = append(orderResponse.OrderLines, orderLineResponse)
		}

		orderListResponse.Orders = append(orderListResponse.Orders, orderResponse)
	}

	return orderListResponse, nil
}
