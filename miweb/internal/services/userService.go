package services

import (
	"paquetes/internal/models"
	"paquetes/internal/repositories"
)

type UsersService struct {
	usersRepository *repositories.UsersRepository
}

func NewUsersService(usersRepository *repositories.UsersRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}

func (s *UsersService) Create(user *models.User) error {
	// Validar que el usuario tenga todos los campos requeridos

	// Validar que no exista ya un usuario con el mismo correo electrónico

	// Encriptar la contraseña

	// Llamar al repositorio para crear el usuario en la base de datos
	err := s.usersRepository.Create(user)
	if err != nil {
		return err
	}

	return nil
}
