package service

import (
	"errors"
	"service/internal/domain/entities"
	"service/internal/repository"
)

type SaleDetailService interface {
	CreateSaleDetail(detail *entities.SaleDetail) error
	GetSaleDetailByID(id uint) (*entities.SaleDetail, error)
	GetAllBySaleID(saleID uint) ([]entities.SaleDetail, error)
	UpdateSaleDetail(detail *entities.SaleDetail) error
	DeleteSaleDetail(id uint) error
}

type saleDetailServiceImpl struct {
	repo repository.SaleDetailRepository
}

func NewSaleDetailService(repo repository.SaleDetailRepository) SaleDetailService {
	return &saleDetailServiceImpl{repo: repo}
}

func (s *saleDetailServiceImpl) CreateSaleDetail(detail *entities.SaleDetail) error {
	if detail.Quantity <= 0 || detail.UnitPrice <= 0 {
		return errors.New("la cantidad y el precio unitario deben ser positivos")
	}
	return s.repo.Create(detail)
}

func (s *saleDetailServiceImpl) GetSaleDetailByID(id uint) (*entities.SaleDetail, error) {
	return s.repo.GetByID(id)
}

func (s *saleDetailServiceImpl) GetAllBySaleID(saleID uint) ([]entities.SaleDetail, error) {
	return s.repo.GetAllBySaleID(saleID)
}

func (s *saleDetailServiceImpl) UpdateSaleDetail(detail *entities.SaleDetail) error {
	if detail.Quantity <= 0 || detail.UnitPrice <= 0 {
		return errors.New("la cantidad y el precio unitario deben ser positivos")
	}
	return s.repo.Update(detail)
}

func (s *saleDetailServiceImpl) DeleteSaleDetail(id uint) error {
	return s.repo.Delete(id)
}
