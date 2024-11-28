package services

import (
	"product-management-project/models"
	"product-management-project/repositories"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService(productRepo *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) FindAllProduct() ([]models.ProductModel, error) {
	return s.productRepo.FindAllProducts()
}

func (s *ProductService) FindProductById(id int) (models.ProductModel, error) {
	return s.productRepo.FindProductById(id)
}

func (s *ProductService) CreateProduct(product models.ProductModel, userId int) (models.ProductModel, error) {
	return s.productRepo.CreateProduct(product, userId)
}

func (s *ProductService) UpdateProduct(product models.ProductModel, id int) (models.ProductModel, error) {
	return s.productRepo.UpdateProduct(product, id)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.productRepo.DeleteProduct(id)
}
