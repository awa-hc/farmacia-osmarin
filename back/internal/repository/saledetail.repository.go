package repository

import (
	"service/internal/domain/entities"

	"gorm.io/gorm"
)

type SaleDetailRepository interface {
	Create(detail *entities.SaleDetail) error
	GetByID(id uint) (*entities.SaleDetail, error)
	GetAllBySaleID(saleID uint) ([]entities.SaleDetail, error)
	Update(detail *entities.SaleDetail) error
	Delete(id uint) error
	DeleteBySaleID(saleID uint) error // Nuevo método
}

type gormSaleDetailRepository struct {
	db *gorm.DB
}

func NewSaleDetailRepository(db *gorm.DB) SaleDetailRepository {
	return &gormSaleDetailRepository{db: db}
}

func (r *gormSaleDetailRepository) Create(detail *entities.SaleDetail) error {
	return r.db.Create(detail).Error
}

func (r *gormSaleDetailRepository) GetByID(id uint) (*entities.SaleDetail, error) {
	var detail entities.SaleDetail
	err := r.db.First(&detail, id).Error
	return &detail, err
}

func (r *gormSaleDetailRepository) GetAllBySaleID(saleID uint) ([]entities.SaleDetail, error) {
	var details []entities.SaleDetail
	err := r.db.Where("sale_id = ?", saleID).Find(&details).Error
	return details, err
}

func (r *gormSaleDetailRepository) Update(detail *entities.SaleDetail) error {
	return r.db.Save(detail).Error
}

func (r *gormSaleDetailRepository) Delete(id uint) error {
	return r.db.Delete(&entities.SaleDetail{}, id).Error
}

func (r *gormSaleDetailRepository) DeleteBySaleID(saleID uint) error {
	return r.db.Where("sale_id = ?", saleID).Delete(&entities.SaleDetail{}).Error
}
