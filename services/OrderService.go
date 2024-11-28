package services

import (
	"product-management-project/models"
	"product-management-project/repositories"
)

type OrderService struct {
	orderRepo *repositories.OrderRepository
}

func NewOrderService(orderRepo *repositories.OrderRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (s *OrderService) CreateOrder(order models.OrderModel) (models.OrderModel, error) {
	return s.orderRepo.CreateOrder(order)
}

func (s *OrderService) FindAllOrders() ([]models.OrderModel, error) {
	return s.orderRepo.FindAllOrders()
}

func (s *OrderService) FindOrderById(id int) (models.OrderModel, error) {
	return s.orderRepo.FindOrderById(id)
}
