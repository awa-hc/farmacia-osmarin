package repository

import (
	"service/internal/domain/entities"

	"gorm.io/gorm"
)

type PurchaseDetailRepository interface {
	Create(detail *entities.PurchaseDetail) error
	GetByID(id uint) (*entities.PurchaseDetail, error)
	GetAllByPurchaseID(purchaseID uint) ([]entities.PurchaseDetail, error)
	Update(detail *entities.PurchaseDetail) error
	Delete(id uint) error
	DeleteByPurchaseID(purchaseID uint) error // Nuevo método
}

type gormPurchaseDetailRepository struct {
	db *gorm.DB
}

func NewPurchaseDetailRepository(db *gorm.DB) PurchaseDetailRepository {
	return &gormPurchaseDetailRepository{db: db}
}

func (r *gormPurchaseDetailRepository) Create(detail *entities.PurchaseDetail) error {
	return r.db.Create(detail).Error
}

func (r *gormPurchaseDetailRepository) GetByID(id uint) (*entities.PurchaseDetail, error) {
	var detail entities.PurchaseDetail
	err := r.db.First(&detail, id).Error
	return &detail, err
}

func (r *gormPurchaseDetailRepository) GetAllByPurchaseID(purchaseID uint) ([]entities.PurchaseDetail, error) {
	var details []entities.PurchaseDetail
	err := r.db.Where("purchase_id = ?", purchaseID).Find(&details).Error
	return details, err
}

func (r *gormPurchaseDetailRepository) Update(detail *entities.PurchaseDetail) error {
	return r.db.Save(detail).Error
}

func (r *gormPurchaseDetailRepository) Delete(id uint) error {
	return r.db.Delete(&entities.PurchaseDetail{}, id).Error
}

func (r *gormPurchaseDetailRepository) DeleteByPurchaseID(purchaseID uint) error {
	return r.db.Where("purchase_id = ?", purchaseID).Delete(&entities.PurchaseDetail{}).Error
}
