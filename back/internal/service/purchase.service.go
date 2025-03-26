package service

import (
	"errors"
	"service/internal/domain/entities"
	"service/internal/repository"
)

type PurchaseService interface {
	CreatePurchase(purchase *entities.Purchase) error
	GetPurchaseByID(id uint) (*entities.Purchase, error)
	GetAllPurchases() ([]entities.Purchase, error)
	UpdatePurchase(purchase *entities.Purchase) error
	DeletePurchase(id uint) error
}

type purchaseServiceImpl struct {
	repo       repository.PurchaseRepository
	detailRepo repository.PurchaseDetailRepository
}

func NewPurchaseService(repo repository.PurchaseRepository, detailRepo repository.PurchaseDetailRepository) PurchaseService {
	return &purchaseServiceImpl{
		repo:       repo,
		detailRepo: detailRepo,
	}
}

func (s *purchaseServiceImpl) CreatePurchase(purchase *entities.Purchase) error {
	if len(purchase.Details) == 0 {
		return errors.New("una compra debe tener al menos un detalle")
	}

	// Validar cada detalle antes de crear la compra
	for _, detail := range purchase.Details {
		if detail.Quantity <= 0 || detail.UnitPrice <= 0 {
			return errors.New("la cantidad y el precio unitario de los detalles deben ser positivos")
		}
	}

	// Crear la compra en la base de datos
	if err := s.repo.Create(purchase); err != nil {
		return err
	}

	// Guardar los detalles asociados a la compra
	for _, detail := range purchase.Details {
		detail.PurchaseID = purchase.ID // Asociar el detalle con la compra creada
		if err := s.detailRepo.Create(&detail); err != nil {
			return err
		}
	}

	return nil
}

func (s *purchaseServiceImpl) GetPurchaseByID(id uint) (*entities.Purchase, error) {
	return s.repo.GetByID(id)
}

func (s *purchaseServiceImpl) GetAllPurchases() ([]entities.Purchase, error) {
	return s.repo.GetAll()
}

func (s *purchaseServiceImpl) UpdatePurchase(purchase *entities.Purchase) error {
	if len(purchase.Details) == 0 {
		return errors.New("una compra debe tener al menos un detalle")
	}

	// Actualizar la compra en la base de datos
	if err := s.repo.Update(purchase); err != nil {
		return err
	}

	// Eliminar los detalles antiguos y crear los nuevos
	if err := s.detailRepo.DeleteByPurchaseID(purchase.ID); err != nil {
		return err
	}

	for _, detail := range purchase.Details {
		detail.PurchaseID = purchase.ID
		if err := s.detailRepo.Create(&detail); err != nil {
			return err
		}
	}

	return nil
}

func (s *purchaseServiceImpl) DeletePurchase(id uint) error {
	// Eliminar los detalles asociados a la compra
	if err := s.detailRepo.DeleteByPurchaseID(id); err != nil {
		return err
	}

	// Eliminar la compra
	return s.repo.Delete(id)
}
