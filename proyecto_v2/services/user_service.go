package services

import (
	"proyecto/models"
	"proyecto/repositories"
)

type UserService interface {
    GetAllUsers() ([]models.User, error)
    GetUserByID(id int) (*models.User, error)
    CreateUser(user *models.User) error
    UpdateUser(user *models.User) error
    DeleteUser(id int) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}

func (us *userService) GetAllUsers() ([]models.User, error) {
	return us.userRepo.GetAllUsers()
}

func (us *userService) GetUserByID(id int) (*models.User, error) {
	return us.userRepo.GetUserByID(id)
}

func (us *userService) CreateUser(user *models.User) error {
	return us.userRepo.CreateUser(user)
}

func (us *userService) UpdateUser(user *models.User) error {
	return us.userRepo.UpdateUser(user)
}

func (us *userService) DeleteUser(id int) error {
	return us.userRepo.DeleteUser(id)
}
