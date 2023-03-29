package database

import (
	"fmt"
	"proyecto/config"
	"proyecto/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	// Crear la cadena de conexión a la base de datos MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	// Conectar a la base de datos MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrar el modelo de usuario a la base de datos
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	// Migrar el modelo de producto a la base de datos
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		return nil, err
	}

	// Asignar la conexión a la variable global DB
	DB = db

	return db, nil
}
