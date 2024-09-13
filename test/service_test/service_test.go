package test

import (
	"order-service/internal/models/base"
	"order-service/internal/models/dto"
	"order-service/internal/models/entity"
	"order-service/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository
type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Create(order entity.TOrderModel) (entity.TOrderModel, error) {
	args := m.Called(order)
	return args.Get(0).(entity.TOrderModel), args.Error(1)
}

func (m *MockOrderRepository) FindByID(id uint) (entity.TOrderModel, error) {
	args := m.Called(id)
	return args.Get(0).(entity.TOrderModel), args.Error(1)
}

func (m *MockOrderRepository) FindAll() ([]entity.TOrderModel, error) {
	args := m.Called()
	return args.Get(0).([]entity.TOrderModel), args.Error(1)
}

func (m *MockOrderRepository) Update(id uint, order entity.TOrderModel) (entity.TOrderModel, error) {
	args := m.Called(id, order)
	return args.Get(0).(entity.TOrderModel), args.Error(1)
}

func (m *MockOrderRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// Mock OrderLineRepository
type MockOrderLineRepository struct {
	mock.Mock
}

func (m *MockOrderLineRepository) Create(orderLine entity.TOrderLinesModel) (entity.TOrderLinesModel, error) {
	args := m.Called(orderLine)
	return args.Get(0).(entity.TOrderLinesModel), args.Error(1)
}

func (m *MockOrderLineRepository) FindByID(id uint) (entity.TOrderLinesModel, error) {
	args := m.Called(id)
	return args.Get(0).(entity.TOrderLinesModel), args.Error(1)
}

func (m *MockOrderLineRepository) FindAll() ([]entity.TOrderLinesModel, error) {
	args := m.Called()
	return args.Get(0).([]entity.TOrderLinesModel), args.Error(1)
}

func (m *MockOrderLineRepository) Update(id uint, orderLine entity.TOrderLinesModel) (entity.TOrderLinesModel, error) {
	args := m.Called(id, orderLine)
	return args.Get(0).(entity.TOrderLinesModel), args.Error(1)
}

func (m *MockOrderLineRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateOrderWithLines(t *testing.T) {
	mockOrderRepo := new(MockOrderRepository)
	mockOrderLineRepo := new(MockOrderLineRepository)

	orderService := service.NewOrderService(mockOrderRepo, mockOrderLineRepo)

	orderDTO := dto.OrderDTO{
		TableID:   1,
		Total:     100.0,
		Createdby: 1,
		OrderLines: []dto.OrderLineDTO{
			{MenuID: 1, Quantity: 2, Price: 10.0},
		},
	}

	createdOrder := entity.TOrderModel{
		Entity: base.Entity{
			Createdby: 1,
		},
		TOrderEntity: entity.TOrderEntity{
			TableID: 1,
			Total:   100.0,
		},
	}

	mockOrderRepo.On("Create", mock.Anything).Return(createdOrder, nil)
	mockOrderLineRepo.On("Create", mock.Anything).Return(entity.TOrderLinesModel{}, nil)

	order, err := orderService.CreateOrderWithLines(orderDTO)
	assert.NoError(t, err)
	assert.Equal(t, createdOrder.ID, order.ID)
}

func TestGetOrderByID(t *testing.T) {
	mockOrderRepo := new(MockOrderRepository)
	orderService := service.NewOrderService(mockOrderRepo, nil)

	expectedOrder := entity.TOrderModel{
		Entity: base.Entity{
			Createdby: 1,
		},
		TOrderEntity: entity.TOrderEntity{
			TableID: 1,
			Total:   100.0,
		},
	}

	mockOrderRepo.On("FindByID", uint(1)).Return(expectedOrder, nil)

	order, err := orderService.GetOrderByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedOrder.ID, order.ID)
}

func TestGetAllOrders(t *testing.T) {
	mockOrderRepo := new(MockOrderRepository)
	orderService := service.NewOrderService(mockOrderRepo, nil)

	expectedOrders := []entity.TOrderModel{
		entity.TOrderModel{
			Entity: base.Entity{
				Createdby: 1,
			},
			TOrderEntity: entity.TOrderEntity{
				TableID: 1,
				Total:   100.0,
			},
		},
		entity.TOrderModel{
			Entity: base.Entity{
				Createdby: 1,
			},
			TOrderEntity: entity.TOrderEntity{
				TableID: 2,
				Total:   200.0,
			},
		},
	}

	mockOrderRepo.On("FindAll").Return(expectedOrders, nil)

	orders, err := orderService.GetAllOrders()
	assert.NoError(t, err)
	assert.Equal(t, expectedOrders, orders)
}
