package repositories

import (
	"proyecto/database"
	"proyecto/models"

	"gorm.io/gorm"
)

// UserRepository interface defines methods to interact with user data source
type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (ur *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, uint(id)).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) CreateUser(user *models.User) error {
	if err := database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) UpdateUser(user *models.User) error {
	if err := database.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) DeleteUser(id int) error {
	user := models.User{Model: gorm.Model{ID: uint(id)}}
	if err := database.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
