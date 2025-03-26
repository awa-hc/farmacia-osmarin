package entities

import (
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	CustomerName string       `gorm:"size:255;not null" json:"customer_name"` // Nombre del cliente
	SaleDate     time.Time    `gorm:"not null" json:"sale_date"`              // Fecha de la venta
	TotalAmount  float64      `gorm:"not null" json:"total_amount"`           // Monto total de la venta
	Status       string       `gorm:"size:50;not null" json:"status"`         // Estado de la venta (ej. Pendiente, Completada)
	Details      []SaleDetail `gorm:"foreignKey:SaleID" json:"details"`       // Detalles de la venta
}
