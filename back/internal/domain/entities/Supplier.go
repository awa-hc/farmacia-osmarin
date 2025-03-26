package entities

import (
	"gorm.io/gorm"
)

type Supplier struct {
	gorm.Model
	Name        string `gorm:"size:255;not null" json:"name"` // Nombre del proveedor
	Address     string `gorm:"size:255" json:"address"`       // Dirección del proveedor
	Phone       string `gorm:"size:20" json:"phone"`          // Teléfono del proveedor
	Email       string `gorm:"size:100" json:"email"`         // Email del proveedor
	Description string `gorm:"size:500" json:"description"`   // Descripción adicional
}
