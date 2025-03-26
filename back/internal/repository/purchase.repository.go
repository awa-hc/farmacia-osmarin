package repository

import (
	"service/internal/domain/entities"

	"gorm.io/gorm"
)

type PurchaseRepository interface {
	Create(purchase *entities.Purchase) error
	GetByID(id uint) (*entities.Purchase, error)
	GetAll() ([]entities.Purchase, error)
	Update(purchase *entities.Purchase) error
	Delete(id uint) error
}

type gormPurchaseRepository struct {
	db *gorm.DB
}

func NewPurchaseRepository(db *gorm.DB) PurchaseRepository {
	return &gormPurchaseRepository{db: db}
}

func (r *gormPurchaseRepository) Create(purchase *entities.Purchase) error {
	return r.db.Create(purchase).Error
}

func (r *gormPurchaseRepository) GetByID(id uint) (*entities.Purchase, error) {
	var purchase entities.Purchase
	err := r.db.Preload("Details").First(&purchase, id).Error
	return &purchase, err
}

func (r *gormPurchaseRepository) GetAll() ([]entities.Purchase, error) {
	var purchases []entities.Purchase
	err := r.db.Preload("Details").Find(&purchases).Error
	return purchases, err
}

func (r *gormPurchaseRepository) Update(purchase *entities.Purchase) error {
	return r.db.Save(purchase).Error
}

func (r *gormPurchaseRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Purchase{}, id).Error
}
