package service

import (
	"facade/internal/repository"
	"facade/model"
)

type OrderService struct {
	repository repository.OrderRepository
}

type Order struct {
	ID    string `json:"id"`
	Order string `json:"order"`
}

func (s *OrderService) GetOrder(id string) (*model.Order, error) {
	return s.repository.FindByID(id)
}
