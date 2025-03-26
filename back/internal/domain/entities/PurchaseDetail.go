package entities

import "gorm.io/gorm"

type PurchaseDetail struct {
	gorm.Model
	PurchaseID uint    `gorm:"not null" json:"purchase_id"` // ID de la compra
	ProductID  uint    `gorm:"not null" json:"product_id"`  // ID del producto
	Quantity   int     `gorm:"not null" json:"quantity"`    // Cantidad comprada
	UnitPrice  float64 `gorm:"not null" json:"unit_price"`  // Precio unitario de compra
	Subtotal   float64 `gorm:"not null" json:"subtotal"`    // Subtotal (cantidad * precio unitario)
}
