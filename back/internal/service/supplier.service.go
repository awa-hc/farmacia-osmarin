package service

import (
	"errors"
	"service/internal/domain/entities"
	"service/internal/repository"
)

type SupplierService interface {
	CreateSupplier(supplier *entities.Supplier) error
	GetSupplierByID(id uint) (*entities.Supplier, error)
	GetAllSuppliers() ([]entities.Supplier, error)
	UpdateSupplier(supplier *entities.Supplier) error
	DeleteSupplier(id uint) error
}

type supplierServiceImpl struct {
	repo repository.SupplierRepository
}

func NewSupplierService(repo repository.SupplierRepository) SupplierService {
	return &supplierServiceImpl{repo: repo}
}

func (s *supplierServiceImpl) CreateSupplier(supplier *entities.Supplier) error {
	if supplier.Name == "" || supplier.Phone == "" {
		return errors.New("el nombre y el teléfono son obligatorios")
	}
	return s.repo.Create(supplier)
}

func (s *supplierServiceImpl) GetSupplierByID(id uint) (*entities.Supplier, error) {
	return s.repo.GetByID(id)
}

func (s *supplierServiceImpl) GetAllSuppliers() ([]entities.Supplier, error) {
	return s.repo.GetAll()
}

func (s *supplierServiceImpl) UpdateSupplier(supplier *entities.Supplier) error {
	if supplier.Name == "" || supplier.Phone == "" {
		return errors.New("el nombre y el teléfono son obligatorios")
	}
	return s.repo.Update(supplier)
}

func (s *supplierServiceImpl) DeleteSupplier(id uint) error {
	return s.repo.Delete(id)
}
