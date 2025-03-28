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
	repo        repository.SaleRepository
	detailRepo  repository.SaleDetailRepository
	productRepo repository.ProductRepository // Repositorio de productos para actualizar el stock
}

func NewSaleService(
	repo repository.SaleRepository,
	detailRepo repository.SaleDetailRepository,
	productRepo repository.ProductRepository, // Agregamos el repositorio de productos
) SaleService {
	return &saleServiceImpl{
		repo:        repo,
		detailRepo:  detailRepo,
		productRepo: productRepo,
	}
}

// Crear una nueva venta
func (s *saleServiceImpl) CreateSale(sale *entities.Sale) error {
	// Iniciar transacción
	tx := s.repo.BeginTransaction()
	if tx == nil {
		return errors.New("no se pudo iniciar la transacción")
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // Revertir transacción en caso de pánico
		}
	}()

	if len(sale.Details) == 0 {
		tx.Rollback()
		return errors.New("una venta debe tener al menos un detalle")
	}

	totalAmount := 0.0
	for _, detail := range sale.Details {
		if detail.Quantity <= 0 || detail.UnitPrice <= 0 {
			tx.Rollback()
			return errors.New("la cantidad y el precio unitario de los detalles deben ser positivos")
		}

		product, err := s.productRepo.GetByID(detail.ProductID)
		if err != nil {
			tx.Rollback()
			return errors.New("producto no encontrado: " + err.Error())
		}

		if product.Stock < detail.Quantity {
			tx.Rollback()
			return errors.New("stock insuficiente para el producto: " + product.Name)
		}

		detail.Subtotal = float64(detail.Quantity) * product.SellingPrice
		totalAmount += detail.Subtotal
	}

	sale.TotalAmount = totalAmount

	// Crear la venta en la base de datos
	if err := s.repo.Create(tx, sale); err != nil {
		tx.Rollback()
		return err
	}

	for _, detail := range sale.Details {
		detail.SaleID = sale.ID

		if err := s.detailRepo.Create(tx, &detail); err != nil {
			tx.Rollback()
			return err
		}

		product, _ := s.productRepo.GetByID(detail.ProductID)
		product.Stock -= detail.Quantity
		if err := s.productRepo.Update(tx, product); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Confirmar transacción
	return tx.Commit().Error
}

// Obtener una venta por ID
func (s *saleServiceImpl) GetSaleByID(id uint) (*entities.Sale, error) {
	return s.repo.GetByID(id)
}

// Obtener todas las ventas
func (s *saleServiceImpl) GetAllSales() ([]entities.Sale, error) {
	return s.repo.GetAll()
}

// Actualizar una venta
func (s *saleServiceImpl) UpdateSale(sale *entities.Sale) error {
	if len(sale.Details) == 0 {
		return errors.New("una venta debe tener al menos un detalle")
	}

	// Eliminar los detalles antiguos
	if err := s.detailRepo.DeleteBySaleID(sale.ID); err != nil {
		return err
	}

	// Validar y crear los nuevos detalles
	totalAmount := 0.0
	for _, detail := range sale.Details {
		if detail.Quantity <= 0 || detail.UnitPrice <= 0 {
			return errors.New("la cantidad y el precio unitario de los detalles deben ser positivos")
		}

		// Obtener el producto asociado
		product, err := s.productRepo.GetByID(detail.ProductID)
		if err != nil {
			return errors.New("producto no encontrado: " + err.Error())
		}

		// Verificar que haya suficiente stock
		if product.Stock < detail.Quantity {
			return errors.New("stock insuficiente para el producto: " + product.Name)
		}

		// Calcular el subtotal
		detail.Subtotal = float64(detail.Quantity) * product.SellingPrice
		totalAmount += detail.Subtotal
	}

	// Asignar el total a la venta
	sale.TotalAmount = totalAmount

	// Actualizar la venta en la base de datos
	if err := s.repo.Update(sale); err != nil {
		return err
	}

	// Guardar los nuevos detalles y actualizar el stock
	for _, detail := range sale.Details {
		// Asociar el detalle con la venta
		detail.SaleID = sale.ID

		// Guardar el detalle en la base de datos
		if err := s.detailRepo.Create(&detail); err != nil {
			return err
		}

		// Actualizar el stock del producto
		product, _ := s.productRepo.GetByID(detail.ProductID)
		product.Stock -= detail.Quantity
		if err := s.productRepo.Update(product); err != nil {
			return err
		}
	}

	return nil
}

// Eliminar una venta
func (s *saleServiceImpl) DeleteSale(id uint) error {
	// Obtener la venta para recuperar los detalles
	sale, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("venta no encontrada: " + err.Error())
	}

	// Restaurar el stock de los productos involucrados
	for _, detail := range sale.Details {
		product, _ := s.productRepo.GetByID(detail.ProductID)
		product.Stock += detail.Quantity
		if err := s.productRepo.Update(product); err != nil {
			return err
		}
	}

	// Eliminar los detalles asociados a la venta
	if err := s.detailRepo.DeleteBySaleID(id); err != nil {
		return err
	}

	// Eliminar la venta
	return s.repo.Delete(id)
}
