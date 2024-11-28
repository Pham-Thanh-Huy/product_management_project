package services

import (
	"product-management-project/models"
	"product-management-project/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetAllUsers() ([]models.UserModel, error) {
	return s.userRepo.FindAll()
}

func (s *UserService) GetUserById(id int) (models.UserModel, error) {
	return s.userRepo.FindById(id)
}

func (s *UserService) CreateUser(user models.UserModel) (models.UserModel, error) {
	return s.userRepo.Create(user)
}

func (s *UserService) UpdateUser(user models.UserModel) error {
	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.userRepo.Delete(id)
}
