package repository

import (
	"gorm.io/gorm"
)

type BaseRepository struct {
	DB *gorm.DB
}

func (r *BaseRepository) WithTransaction(fn func(db *gorm.DB) error) error {
	tx := r.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
