package service

import (
	"errors"
	"go-sqlite-crud-project/model"
	"go-sqlite-crud-project/repository"
)

type ProductService struct {
	ProductRepo *repository.ProductRepository
}

func NewProductService(ProductRepo *repository.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: ProductRepo}
}

// Add a new product
func (service *ProductService) CreateProduct(product *model.Product) (*model.Product, error) {
	if product.Name == "" || product.Description == "" || product.Price <= 0 || product.Stock < 0 || product.CategoryID <= 0 {
		return nil, errors.New("invalid product data")
	}

	createdProduct, err := service.ProductRepo.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return createdProduct, nil
}

// Fetch product by ID
func (service *ProductService) GetProduct(id int) (*model.Product, error) {
	if id <= 0 {
		return nil, errors.New("invalid product ID")
	}

	product, err := service.ProductRepo.GetProduct(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

// Fetch all products with pagination
func (service *ProductService) GetAllProducts(page, limit int) ([]model.Product, error) {
	if page <= 0 || limit <= 0 {
		return nil, errors.New("invalid pagination parameters")
	}

	products, err := service.ProductRepo.GetAllProducts(page, limit)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// Update product details
func (service *ProductService) UpdateProduct(product *model.Product) (*model.Product, error) {
	if product.ID <= 0 || product.Name == "" || product.Description == "" || product.Price <= 0 || product.Stock < 0 || product.CategoryID <= 0 {
		return nil, errors.New("invalid product data")
	}

	existingProduct, err := service.ProductRepo.GetProduct(product.ID)
	if err != nil {
		return nil, err
	}
	if existingProduct == nil {
		return nil, errors.New("product not found")
	}

	updatedProduct, err := service.ProductRepo.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

// Delete a product by ID
func (service *ProductService) DeleteProduct(id int) error {
	if id <= 0 {
		return errors.New("invalid product ID")
	}

	existingProduct, err := service.ProductRepo.GetProduct(id)
	if err != nil {
		return err
	}
	if existingProduct == nil {
		return errors.New("product not found")
	}

	return service.ProductRepo.DeleteProduct(id)
}

// package service

// import (
// 	"go-sqlite-crud-project/model"
// 	"go-sqlite-crud-project/repository"
// )

// type ProductService struct {
// 	ProductRepo *repository.ProductRepository
// }

// func NewProductService(productRepo *repository.ProductRepository) *ProductService {
// 	return &ProductService{ProductRepo: productRepo}
// }

// func (service *ProductService) CreateProduct(product *model.Product) (*model.Product, error) {
// 	// Business logic goes here
// 	return service.ProductRepo.CreateProduct(product)
// }

// func (service *ProductService) GetProduct(id int) (*model.Product, error) {
// 	return service.ProductRepo.GetProduct(id)
// }

// func (service *ProductService) GetAllProducts() ([]model.Product, error) {
// 	return service.ProductRepo.GetAllProducts()
// }

// func (service *ProductService) UpdateProduct(product *model.Product) (*model.Product, error) {
// 	return service.ProductRepo.UpdateProduct(product)
// }

// func (service *ProductService) DeleteProduct(id int) error {
// 	return service.ProductRepo.DeleteProduct(id)
// }
