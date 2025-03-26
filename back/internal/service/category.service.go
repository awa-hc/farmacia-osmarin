package service

import (
	"errors"
	"service/internal/domain/entities"
	"service/internal/repository"
)

type CategoryService interface {
	CreateCategory(category *entities.Category) error
	GetCategoryByID(id uint) (*entities.Category, error)
	GetAllCategories() ([]entities.Category, error)
	UpdateCategory(category *entities.Category) error
	DeleteCategory(id uint) error
}

type categoryServiceImpl struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryServiceImpl{repo: repo}
}

func (s *categoryServiceImpl) CreateCategory(category *entities.Category) error {
	if category.Name == "" {
		return errors.New("el nombre de la categoría es obligatorio")
	}
	return s.repo.Create(category)
}

func (s *categoryServiceImpl) GetCategoryByID(id uint) (*entities.Category, error) {
	return s.repo.GetByID(id)
}

func (s *categoryServiceImpl) GetAllCategories() ([]entities.Category, error) {
	return s.repo.GetAll()
}

func (s *categoryServiceImpl) UpdateCategory(category *entities.Category) error {
	if category.Name == "" {
		return errors.New("el nombre de la categoría es obligatorio")
	}
	return s.repo.Update(category)
}

func (s *categoryServiceImpl) DeleteCategory(id uint) error {
	return s.repo.Delete(id)
}
