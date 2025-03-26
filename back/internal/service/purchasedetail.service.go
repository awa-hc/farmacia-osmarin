package service

import (
	"errors"
	"service/internal/domain/entities"
	"service/internal/repository"
)

type PurchaseDetailService interface {
	CreatePurchaseDetail(detail *entities.PurchaseDetail) error
	GetPurchaseDetailByID(id uint) (*entities.PurchaseDetail, error)
	GetAllByPurchaseID(purchaseID uint) ([]entities.PurchaseDetail, error)
	UpdatePurchaseDetail(detail *entities.PurchaseDetail) error
	DeletePurchaseDetail(id uint) error
}

type purchaseDetailServiceImpl struct {
	repo repository.PurchaseDetailRepository
}

func NewPurchaseDetailService(repo repository.PurchaseDetailRepository) PurchaseDetailService {
	return &purchaseDetailServiceImpl{repo: repo}
}

func (s *purchaseDetailServiceImpl) CreatePurchaseDetail(detail *entities.PurchaseDetail) error {
	if detail.Quantity <= 0 || detail.UnitPrice <= 0 {
		return errors.New("la cantidad y el precio unitario deben ser positivos")
	}
	return s.repo.Create(detail)
}

func (s *purchaseDetailServiceImpl) GetPurchaseDetailByID(id uint) (*entities.PurchaseDetail, error) {
	return s.repo.GetByID(id)
}

func (s *purchaseDetailServiceImpl) GetAllByPurchaseID(purchaseID uint) ([]entities.PurchaseDetail, error) {
	return s.repo.GetAllByPurchaseID(purchaseID)
}

func (s *purchaseDetailServiceImpl) UpdatePurchaseDetail(detail *entities.PurchaseDetail) error {
	if detail.Quantity <= 0 || detail.UnitPrice <= 0 {
		return errors.New("la cantidad y el precio unitario deben ser positivos")
	}
	return s.repo.Update(detail)
}

func (s *purchaseDetailServiceImpl) DeletePurchaseDetail(id uint) error {
	return s.repo.Delete(id)
}
