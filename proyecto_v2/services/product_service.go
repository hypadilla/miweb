package services

import (
	"proyecto/models"
	"proyecto/repositories"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id int) (*models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id int) error
}

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{productRepo}
}

func (us *productService) GetAllProducts() ([]models.Product, error) {
	return us.productRepo.GetAllProducts()
}

func (us *productService) GetProductByID(id int) (*models.Product, error) {
	return us.productRepo.GetProductByID(id)
}

func (us *productService) CreateProduct(product *models.Product) error {
	return us.productRepo.CreateProduct(product)
}

func (us *productService) UpdateProduct(product *models.Product) error {
	return us.productRepo.UpdateProduct(product)
}

func (us *productService) DeleteProduct(id int) error {
	return us.productRepo.DeleteProduct(id)
}
