package repository

import (
	"order-service/internal/models/entity"
	"order-service/internal/pkg/logging"

	"gorm.io/gorm"
)

type OrderLineRepository interface {
	Create(orderLine entity.TOrderLinesModel) (entity.TOrderLinesModel, error)
	FindByID(id uint) (entity.TOrderLinesModel, error)
	FindAll() ([]entity.TOrderLinesModel, error)
	Update(id uint, orderLine entity.TOrderLinesModel) (entity.TOrderLinesModel, error)
	Delete(id uint) error
}

type orderLineRepository struct {
	db *gorm.DB
}

func NewOrderLineRepository(db *gorm.DB) OrderLineRepository {
	return &orderLineRepository{db: db}
}

func (o *orderLineRepository) Create(orderLine entity.TOrderLinesModel) (entity.TOrderLinesModel, error) {
	err := o.db.Create(&orderLine).Error // Menggunakan pointer pada orderLine
	if err != nil {
		logging.Log.Error("Failed to create order line:", err)
		return orderLine, err
	}
	return orderLine, nil
}

func (o *orderLineRepository) FindByID(id uint) (entity.TOrderLinesModel, error) {
	var orderLine entity.TOrderLinesModel
	err := o.db.Where("id = ?", id).First(&orderLine).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logging.Log.Warnf("Order line with ID %d not found", id)
		} else {
			logging.Log.Error("Failed to find order line by ID:", err)
		}
		return orderLine, err
	}
	return orderLine, nil
}

func (o *orderLineRepository) FindAll() ([]entity.TOrderLinesModel, error) {
	var orderLines []entity.TOrderLinesModel
	err := o.db.Find(&orderLines).Error
	if err != nil {
		logging.Log.Error("Failed to find all order lines:", err)
		return orderLines, err
	}
	return orderLines, nil
}

func (o *orderLineRepository) Update(id uint, orderLine entity.TOrderLinesModel) (entity.TOrderLinesModel, error) {
	existingRecord, err := o.FindByID(id)
	if err != nil {
		logging.Log.Error("Failed to find order line for update:", err)
		return orderLine, err
	}

	updatedColumn := map[string]interface{}{
		"order_id":  orderLine.OrderID,
		"menu_id":   orderLine.MenuID,
		"quantity":  orderLine.Quantity,
		"price":     orderLine.Price,
		"sub_total": orderLine.SubTotal,
	}

	err = o.db.Model(&existingRecord).Updates(updatedColumn).Error
	if err != nil {
		logging.Log.Error("Failed to update order line:", err)
		return existingRecord, err // Mengembalikan record yang telah diperbarui
	}
	return existingRecord, nil // Mengembalikan existingRecord, bukan orderLine
}

func (o *orderLineRepository) Delete(id uint) error {
	err := o.db.Where("id = ?", id).Delete(&entity.TOrderLinesModel{}).Error
	if err != nil {
		logging.Log.Error("Failed to delete order line:", err)
		return err
	}
	return nil
}
