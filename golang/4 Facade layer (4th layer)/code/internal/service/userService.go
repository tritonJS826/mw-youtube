package service

import (
	"facade/internal/repository"
	"facade/model"
)

type UserService struct {
	repository repository.UserRepository
}

func (s *UserService) GetUser(id string) (*model.User, error) {
	return s.repository.FindByID(id)
}
