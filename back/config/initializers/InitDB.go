package initializers

import (
	"fmt"
	"os"
	"service/internal/domain/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	if err := db.AutoMigrate(
		&entities.Category{},
		&entities.Product{},
		&entities.Purchase{},
		&entities.PurchaseDetail{},
		&entities.Sale{},
		&entities.SaleDetail{},
		&entities.Supplier{},
	); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil

}
