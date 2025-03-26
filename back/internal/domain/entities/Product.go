package entities

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name             string    `gorm:"size:255;not null" json:"name"`              // Nombre del medicamento
	ActiveIngredient string    `gorm:"size:255;not null" json:"active_ingredient"` // Principio activo
	Presentation     string    `gorm:"size:255;not null" json:"presentation"`      // Presentación (ej. Tabletas, Jarabe, etc.)
	ExpiryDate       time.Time `gorm:"not null" json:"expiry_date"`                // Fecha de vencimiento del medicamento
	Industry         string    `gorm:"size:255;not null" json:"industry"`          // Industria farmacéutica que lo fabrica
	EntryPrice       float64   `gorm:"not null" json:"entry_price"`                // Precio de entrada (costo)
	SellingPrice     float64   `gorm:"not null" json:"selling_price"`              // Precio al comprador (precio de venta)
	Stock            int       `gorm:"not null" json:"stock"`                      // Cantidad disponible en stock
	Category         string    `gorm:"size:100" json:"category"`                   // Categoría (ej. Analgésicos, Antibióticos, etc.)
}
