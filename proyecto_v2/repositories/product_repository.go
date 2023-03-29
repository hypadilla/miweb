package repositories

import (
	"proyecto/database"
	"proyecto/models"

	"gorm.io/gorm"
)

// ProductRepository interface defines methods to interact with product data source
type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id int) (*models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id int) error
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (ur *productRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := database.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (ur *productRepository) GetProductByID(id int) (*models.Product, error) {
	var product models.Product
	if err := database.DB.First(&product, uint(id)).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (ur *productRepository) CreateProduct(product *models.Product) error {
	if err := database.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (ur *productRepository) UpdateProduct(product *models.Product) error {
	if err := database.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}

func (ur *productRepository) DeleteProduct(id int) error {
	product := models.Product{Model: gorm.Model{ID: uint(id)}}
	if err := database.DB.Delete(&product).Error; err != nil {
		return err
	}
	return nil
}
