package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `gorm:"size:100;not null;unique" json:"name"` // Nombre de la categoría
	Description string `gorm:"size:500" json:"description"`          // Descripción de la categoría
}
