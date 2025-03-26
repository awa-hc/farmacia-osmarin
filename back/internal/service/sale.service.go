package service

import (
	"errors"
	"service/internal/domain/entities"
	"service/internal/repository"
)

type SaleService interface {
	CreateSale(sale *entities.Sale) error
	GetSaleByID(id uint) (*entities.Sale, error)
	GetAllSales() ([]entities.Sale, error)
	UpdateSale(sale *entities.Sale) error
	DeleteSale(id uint) error
}

type saleServiceImpl struct {
	repo       repository.SaleRepository
	detailRepo repository.SaleDetailRepository
}

func NewSaleService(repo repository.SaleRepository, detailRepo repository.SaleDetailRepository) SaleService {
	return &saleServiceImpl{
		repo:       repo,
		detailRepo: detailRepo,
	}
}

func (s *saleServiceImpl) CreateSale(sale *entities.Sale) error {
	if len(sale.Details) == 0 {
		return errors.New("una venta debe tener al menos un detalle")
	}

	// Validar cada detalle antes de crear la venta
	for _, detail := range sale.Details {
		if detail.Quantity <= 0 || detail.UnitPrice <= 0 {
			return errors.New("la cantidad y el precio unitario de los detalles deben ser positivos")
		}
	}

	// Crear la venta en la base de datos
	if err := s.repo.Create(sale); err != nil {
		return err
	}

	// Guardar los detalles asociados a la venta
	for _, detail := range sale.Details {
		detail.SaleID = sale.ID // Asociar el detalle con la venta creada
		if err := s.detailRepo.Create(&detail); err != nil {
			return err
		}
	}

	return nil
}

func (s *saleServiceImpl) GetSaleByID(id uint) (*entities.Sale, error) {
	return s.repo.GetByID(id)
}

func (s *saleServiceImpl) GetAllSales() ([]entities.Sale, error) {
	return s.repo.GetAll()
}

func (s *saleServiceImpl) UpdateSale(sale *entities.Sale) error {
	if len(sale.Details) == 0 {
		return errors.New("una venta debe tener al menos un detalle")
	}

	// Actualizar la venta en la base de datos
	if err := s.repo.Update(sale); err != nil {
		return err
	}

	// Eliminar los detalles antiguos y crear los nuevos
	if err := s.detailRepo.DeleteBySaleID(sale.ID); err != nil {
		return err
	}

	for _, detail := range sale.Details {
		detail.SaleID = sale.ID
		if err := s.detailRepo.Create(&detail); err != nil {
			return err
		}
	}

	return nil
}

func (s *saleServiceImpl) DeleteSale(id uint) error {
	// Eliminar los detalles asociados a la venta
	if err := s.detailRepo.DeleteBySaleID(id); err != nil {
		return err
	}

	// Eliminar la venta
	return s.repo.Delete(id)
}
