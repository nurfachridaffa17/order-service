package repository

import (
	"order-service/internal/models/entity"
	"order-service/internal/pkg/logging"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order entity.TOrderModel) (entity.TOrderModel, error)
	FindByID(id uint) (entity.TOrderModel, error)
	FindAll() ([]entity.TOrderModel, error)
	Update(id uint, order entity.TOrderModel) (entity.TOrderModel, error)
	Delete(id uint) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (o *orderRepository) Create(order entity.TOrderModel) (entity.TOrderModel, error) {
	err := o.db.Create(&order).Error // Use pointer for the entity
	if err != nil {
		logging.Log.Error("Failed to create order:", err)
		return order, err
	}
	return order, nil
}

func (o *orderRepository) FindByID(id uint) (entity.TOrderModel, error) {
	var order entity.TOrderModel
	err := o.db.Preload("OrderLines").Where("id = ?", id).First(&order).Error
	if err != nil {
		logging.Log.Error("Error finding order:", err)
		return order, err
	}
	return order, nil
}

func (o *orderRepository) FindAll() ([]entity.TOrderModel, error) {
	var orders []entity.TOrderModel
	err := o.db.Find(&orders).Error
	if err != nil {
		logging.Log.Error("Error finding all orders:", err)
		return orders, err
	}
	return orders, nil
}

func (o *orderRepository) Update(id uint, order entity.TOrderModel) (entity.TOrderModel, error) {
	existingRecord, err := o.FindByID(id)
	if err != nil {
		logging.Log.Error("Failed to find order for update:", err)
		return order, err
	}

	updatedColumns := map[string]interface{}{
		"table_id": order.TableID,
		"total":    order.Total,
	}

	err = o.db.Model(&existingRecord).Updates(updatedColumns).Error
	if err != nil {
		logging.Log.Error("Failed to update order:", err)
		return existingRecord, err
	}
	return existingRecord, nil // Return the updated record
}

func (o *orderRepository) Delete(id uint) error {
	err := o.db.Where("id = ?", id).Delete(&entity.TOrderModel{}).Error
	if err != nil {
		logging.Log.Error("Failed to delete order:", err)
		return err
	}
	return nil
}
