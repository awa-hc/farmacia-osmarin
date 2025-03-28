package repository

import (
	"gorm.io/gorm"
)

type BaseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

func (r *BaseRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}

func (r *BaseRepository) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *BaseRepository) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}
