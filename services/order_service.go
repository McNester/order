package services

import (
	"orders/models"
	"orders/repositories"
)

type OrderService struct {
	repo *repositories.OrderRepo
}

func NewOrderService() *OrderService {
	return &OrderService{repo: repositories.NewOrderRepo()}
}

func (s *OrderService) SaveOrder(order *models.Order) (*models.Order, error) {
	return s.repo.SaveOrder(order)
}

func (s *OrderService) GetOrder(id uint64) (*models.Order, error) {
	return s.repo.GetOrder(id)
}

func (s *OrderService) UpdateOrder(id uint64, order *models.Order) (*models.Order, error) {
	return s.repo.UpdateOrder(id, order)
}

func (s *OrderService) DeleteOrder(id uint64) error {
	return s.repo.DeleteOrder(id)
}

func (s *OrderService) ListOrder() ([]models.Order, error) {
	return s.repo.ListOrders()
}
