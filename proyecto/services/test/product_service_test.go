package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"proyecto/models"
	"proyecto/repositories"
	"proyecto/services"
)

func TestGetAllProducts(t *testing.T) {
	productRepo := repositories.NewMockProductRepository([]models.Product{
		{Name: "Alice", Description: "alice@example.com", Price: 123},
		{Name: "Bob", Description: "bob@example.com", Price: 123},
	})
	productService := services.NewProductService(productRepo)

	products, err := productService.GetAllProducts()
	assert.NoError(t, err)
	assert.Len(t, products, 2)
}

func TestGetProductByID(t *testing.T) {
	productRepo := repositories.NewMockProductRepository([]models.Product{
		{Model: gorm.Model{ID: 1}, Name: "Alice", Description: "alice@example.com", Price: 123},
		{Model: gorm.Model{ID: 1}, Name: "Bob", Description: "bob@example.com", Price: 123},
	})
	productService := services.NewProductService(productRepo)

	product, err := productService.GetProductByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "Alice", product.Name)

	product, err = productService.GetProductByID(3)
	assert.ErrorIs(t, err, repositories.ErrRecordNotFound)
	assert.Nil(t, product)
}

func TestCreateProduct(t *testing.T) {
	productRepo := repositories.NewMockProductRepository([]models.Product{})
	productService := services.NewProductService(productRepo)

	err := productService.CreateProduct(&models.Product{Name: "Alice", Description: "alice@example.com", Price: 123})
	assert.NoError(t, err)

	products, _ := productService.GetAllProducts()
	assert.Len(t, products, 1)
}

func TestUpdateProduct(t *testing.T) {
	productRepo := repositories.NewMockProductRepository([]models.Product{
		{Model: gorm.Model{ID: 1}, Name: "Alice", Description: "alice@example.com", Price: 13},
	})
	productService := services.NewProductService(productRepo)

	err := productService.UpdateProduct(&models.Product{Model: gorm.Model{ID: 1}, Name: "Alice", Description: "alice@example.com", Price: 123})
	assert.NoError(t, err)

	product, _ := productService.GetProductByID(1)
	assert.Equal(t, 123.0, product.Price)
}

func TestDeleteProduct(t *testing.T) {
	productRepo := repositories.NewMockProductRepository([]models.Product{
		{Model: gorm.Model{ID: 1}, Name: "Alice", Description: "alice@example.com", Price: 123},
		{Model: gorm.Model{ID: 1}, Name: "Bob", Description: "bob@example.com", Price: 123},
	})
	productService := services.NewProductService(productRepo)

	err := productService.DeleteProduct(1)
	assert.NoError(t, err)

	products, _ := productService.GetAllProducts()
	assert.Len(t, products, 1)

	err = productService.DeleteProduct(3)
	assert.ErrorIs(t, err, repositories.ErrRecordNotFound)
}
