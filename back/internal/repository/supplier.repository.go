package repository

import (
	"service/internal/domain/entities"

	"gorm.io/gorm"
)

type SupplierRepository interface {
	Create(supplier *entities.Supplier) error
	GetByID(id uint) (*entities.Supplier, error)
	GetAll() ([]entities.Supplier, error)
	Update(supplier *entities.Supplier) error
	Delete(id uint) error
}

type gormSupplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &gormSupplierRepository{db: db}
}

func (r *gormSupplierRepository) Create(supplier *entities.Supplier) error {
	return r.db.Create(supplier).Error
}

func (r *gormSupplierRepository) GetByID(id uint) (*entities.Supplier, error) {
	var supplier entities.Supplier
	err := r.db.First(&supplier, id).Error
	return &supplier, err
}

func (r *gormSupplierRepository) GetAll() ([]entities.Supplier, error) {
	var suppliers []entities.Supplier
	err := r.db.Find(&suppliers).Error
	return suppliers, err
}

func (r *gormSupplierRepository) Update(supplier *entities.Supplier) error {
	return r.db.Save(supplier).Error
}

func (r *gormSupplierRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Supplier{}, id).Error
}
