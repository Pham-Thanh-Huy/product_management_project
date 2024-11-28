package repositories

import (
	"gorm.io/gorm"
	"product-management-project/models"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order models.OrderModel) (models.OrderModel, error) {
	err := r.db.Create(&order).Error
	return order, err
}

func (r *OrderRepository) FindAllOrders() ([]models.OrderModel, error) {
	var orders []models.OrderModel
	err := r.db.Preload("User").Preload("Items.Product").Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) FindOrderById(id int) (models.OrderModel, error) {
	var order models.OrderModel
	err := r.db.Preload("User").Preload("Items.Product").First(&order, id).Error
	return order, err
}
