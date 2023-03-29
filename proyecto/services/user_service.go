package services

import (
	"fmt"
	"proyecto/models"
	"proyecto/repositories"
	services_cache "proyecto/services/cache"
	"runtime"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
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
	pc, _, _, _ := runtime.Caller(1)
	key := runtime.FuncForPC(pc).Name()
	usuarios := []models.User{}
	if datos, err := services_cache.ObtenerDatosDeCache(key); err == nil {
		usuariosCache := make([]interface{}, len(datos))
		copy(usuariosCache, datos)
		fmt.Println("Usuarios de la caché:", usuariosCache)
		for _, usuario := range usuariosCache {
			usuarios = append(usuarios, usuario.(models.User))
		}
	} else {
		var err error
		usuarios, err = us.userRepo.GetAllUsers()
		if err != nil {
			return nil, err
		}
		usuariosInterfaces := make([]interface{}, len(usuarios))
		for i, v := range usuarios {
			usuariosInterfaces[i] = v
		}
		services_cache.GuardarDatosEnCache(key, usuariosInterfaces)
		fmt.Println("Guarde Usuarios de la caché:", usuariosInterfaces)
	}
	return usuarios, nil
}

func (us *userService) GetUserByID(id int) (*models.User, error) {
	return us.userRepo.GetUserByID(id)
}

func (us *userService) GetUserByEmail(email string) (*models.User, error) {
	return us.userRepo.GetUserByEmail(email)
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
