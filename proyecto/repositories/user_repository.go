package repositories

import (
	"errors"
	"proyecto/database"
	"proyecto/models"

	"gorm.io/gorm"
)

var ErrRecordNotFound = errors.New("record not found")

// UserRepository interface defines methods to interact with user data source
type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (ur *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User = []models.User{}
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

func (ur *userRepository) GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrRecordNotFound
        }
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

// test
func NewMockUserRepository(users []models.User) *MockUserRepository {
	return &MockUserRepository{users: users}
}

type MockUserRepository struct {
	users []models.User
}

func (r *MockUserRepository) GetAllUsers() ([]models.User, error) {
	return r.users, nil
}

func (r *MockUserRepository) GetUserByID(id int) (*models.User, error) {
	for i := range r.users {
		if r.users[i].ID == uint(id) {
			return &r.users[i], nil
		}
	}
	return nil, ErrRecordNotFound
}

func (r *MockUserRepository) GetUserByEmail(email string) (*models.User, error) {
	for i := range r.users {
		if r.users[i].Email == email {
			return &r.users[i], nil
		}
	}
	return nil, ErrRecordNotFound
}

func (r *MockUserRepository) CreateUser(user *models.User) error {
	user.ID = uint(len(r.users) + 1)
	r.users = append(r.users, *user)
	return nil
}

func (r *MockUserRepository) UpdateUser(user *models.User) error {
	for i := range r.users {
		if r.users[i].ID == user.ID {
			r.users[i] = *user
			return nil
		}
	}
	return ErrRecordNotFound
}

func (r *MockUserRepository) DeleteUser(id int) error {
	for i := range r.users {
		if r.users[i].ID == uint(id) {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return ErrRecordNotFound
}
