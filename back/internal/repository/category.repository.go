package repository

import (
	"service/internal/domain/entities"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *entities.Category) error
	GetByID(id uint) (*entities.Category, error)
	GetAll() ([]entities.Category, error)
	Update(category *entities.Category) error
	Delete(id uint) error
}

type gormCategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &gormCategoryRepository{db: db}
}

func (r *gormCategoryRepository) Create(category *entities.Category) error {
	return r.db.Create(category).Error
}

func (r *gormCategoryRepository) GetByID(id uint) (*entities.Category, error) {
	var category entities.Category
	err := r.db.First(&category, id).Error
	return &category, err
}

func (r *gormCategoryRepository) GetAll() ([]entities.Category, error) {
	var categories []entities.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *gormCategoryRepository) Update(category *entities.Category) error {
	return r.db.Save(category).Error
}

func (r *gormCategoryRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Category{}, id).Error
}
