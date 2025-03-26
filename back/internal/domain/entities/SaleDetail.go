package entities

import "gorm.io/gorm"

type SaleDetail struct {
	gorm.Model
	SaleID    uint    `gorm:"not null" json:"sale_id"`    // ID de la venta
	ProductID uint    `gorm:"not null" json:"product_id"` // ID del producto
	Quantity  int     `gorm:"not null" json:"quantity"`   // Cantidad vendida
	UnitPrice float64 `gorm:"not null" json:"unit_price"` // Precio unitario de venta
	Subtotal  float64 `gorm:"not null" json:"subtotal"`   // Subtotal (cantidad * precio unitario)
}
