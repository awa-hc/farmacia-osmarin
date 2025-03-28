package repository

import (
	"service/internal/domain/entities"

	"gorm.io/gorm"
)

type SaleRepository interface {
	Create(tx *gorm.DB, sale *entities.Sale) error
	GetByID(id uint) (*entities.Sale, error)
	GetAll() ([]entities.Sale, error)
	Update(sale *entities.Sale) error
	Delete(id uint) error
	BeginTransaction() *gorm.DB
}

type gormSaleRepository struct {
	BaseRepository
}

func NewSaleRepository(db *gorm.DB) SaleRepository {
	return &gormSaleRepository{
		BaseRepository: *NewBaseRepository(db),
	}
}

func (r *gormSaleRepository) Create(tx *gorm.DB, sale *entities.Sale) error {
	if tx != nil {
		return tx.Create(sale).Error
	}
	return r.db.Create(sale).Error
}

func (r *gormSaleRepository) GetByID(id uint) (*entities.Sale, error) {
	var sale entities.Sale
	err := r.db.Preload("Details").First(&sale, id).Error
	return &sale, err
}

func (r *gormSaleRepository) GetAll() ([]entities.Sale, error) {
	var sales []entities.Sale
	err := r.db.Preload("Details").Find(&sales).Error
	return sales, err
}

func (r *gormSaleRepository) Update(sale *entities.Sale) error {
	return r.db.Save(sale).Error
}

func (r *gormSaleRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Sale{}, id).Error
}
