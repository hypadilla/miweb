package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"proyecto/models"
	"proyecto/repositories"
	"proyecto/services"
)

func TestGetAllUsers(t *testing.T) {
	userRepo := repositories.NewMockUserRepository([]models.User{
		{Name: "Alice", Email: "alice@example.com", Password: "password1"},
		{Name: "Bob", Email: "bob@example.com", Password: "password2"},
	})
	userService := services.NewUserService(userRepo)

	users, err := userService.GetAllUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestGetUserByID(t *testing.T) {
	userRepo := repositories.NewMockUserRepository([]models.User{
		{Model: gorm.Model{ID: 1}, Name: "Alice", Email: "alice@example.com", Password: "password1"},
		{Model: gorm.Model{ID: 1}, Name: "Bob", Email: "bob@example.com", Password: "password2"},
	})
	userService := services.NewUserService(userRepo)

	user, err := userService.GetUserByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "Alice", user.Name)

	user, err = userService.GetUserByID(3)
	assert.ErrorIs(t, err, repositories.ErrRecordNotFound)
	assert.Nil(t, user)
}

func TestCreateUser(t *testing.T) {
	userRepo := repositories.NewMockUserRepository([]models.User{})
	userService := services.NewUserService(userRepo)

	err := userService.CreateUser(&models.User{Name: "Alice", Email: "alice@example.com", Password: "password1"})
	assert.NoError(t, err)

	users, _ := userService.GetAllUsers()
	assert.Len(t, users, 1)
}

func TestUpdateUser(t *testing.T) {
	userRepo := repositories.NewMockUserRepository([]models.User{
		{Model: gorm.Model{ID: 1}, Name: "Alice", Email: "alice@example.com", Password: "password1"},
	})
	userService := services.NewUserService(userRepo)

	err := userService.UpdateUser(&models.User{Model: gorm.Model{ID: 1}, Name: "Alice", Email: "alice@example.com", Password: "password123"})
	assert.NoError(t, err)

	user, _ := userService.GetUserByID(1)
	assert.Equal(t, "password123", user.Password)
}

func TestDeleteUser(t *testing.T) {
	userRepo := repositories.NewMockUserRepository([]models.User{
		{Model: gorm.Model{ID: 1}, Name: "Alice", Email: "alice@example.com", Password: "password1"},
		{Model: gorm.Model{ID: 1}, Name: "Bob", Email: "bob@example.com", Password: "password2"},
	})
	userService := services.NewUserService(userRepo)

	err := userService.DeleteUser(1)
	assert.NoError(t, err)

	users, _ := userService.GetAllUsers()
	assert.Len(t, users, 1)

	err = userService.DeleteUser(3)
	assert.ErrorIs(t, err, repositories.ErrRecordNotFound)
}
