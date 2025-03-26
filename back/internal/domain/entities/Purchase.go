package entities

import (
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	SupplierID   uint             `gorm:"not null" json:"supplier_id"`          // ID del proveedor
	PurchaseDate time.Time        `gorm:"not null" json:"purchase_date"`        // Fecha de la compra
	TotalAmount  float64          `gorm:"not null" json:"total_amount"`         // Monto total de la compra
	Status       string           `gorm:"size:50;not null" json:"status"`       // Estado de la compra (ej. Pendiente, Completada)
	Details      []PurchaseDetail `gorm:"foreignKey:PurchaseID" json:"details"` // Detalles de la compra
}
