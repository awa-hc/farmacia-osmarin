package service

import (
	"errors"
	"service/internal/domain/entities"
	"service/internal/repository"
)

type ProductService interface {
	CreateProduct(product *entities.Product) error
	GetProductByID(id uint) (*entities.Product, error)
	GetAllProducts() ([]entities.Product, error)
	UpdateProduct(product *entities.Product) error
	DeleteProduct(id uint) error
}

type productServiceImpl struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productServiceImpl{repo: repo}
}

func (s *productServiceImpl) CreateProduct(product *entities.Product) error {
	if product.Stock < 0 {
		return errors.New("el stock no puede ser negativo")
	}
	return s.repo.Create(product)
}

func (s *productServiceImpl) GetProductByID(id uint) (*entities.Product, error) {
	return s.repo.GetByID(id)
}

func (s *productServiceImpl) GetAllProducts() ([]entities.Product, error) {
	return s.repo.GetAll()
}

func (s *productServiceImpl) UpdateProduct(product *entities.Product) error {
	if product.Stock < 0 {
		return errors.New("el stock no puede ser negativo")
	}
	return s.repo.Update(product)
}

func (s *productServiceImpl) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}
