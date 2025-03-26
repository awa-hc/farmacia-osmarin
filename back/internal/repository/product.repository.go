package repository

import (
	"service/internal/domain/entities"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *entities.Product) error
	GetByID(id uint) (*entities.Product, error)
	GetAll() ([]entities.Product, error)
	Update(product *entities.Product) error
	Delete(id uint) error
}

type gormProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &gormProductRepository{db: db}
}

func (r *gormProductRepository) Create(product *entities.Product) error {
	return r.db.Create(product).Error
}

func (r *gormProductRepository) GetByID(id uint) (*entities.Product, error) {
	var product entities.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *gormProductRepository) GetAll() ([]entities.Product, error) {
	var products []entities.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *gormProductRepository) Update(product *entities.Product) error {
	return r.db.Save(product).Error
}

func (r *gormProductRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Product{}, id).Error
}
