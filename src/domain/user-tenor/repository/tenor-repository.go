package repository

import (
	"test-kp-golang/src/domain/user-tenor/entity"

	"gorm.io/gorm"
)

type UserTenorRepository struct {
	db *gorm.DB
}

func NewUserTenorRepository(db *gorm.DB) *UserTenorRepository {
	return &UserTenorRepository{
		db: db,
	}
}

func (r *UserTenorRepository) FindByUserId(id int) ([]entity.UserTenor, error) {
	var userTenors []entity.UserTenor

	result := r.db.Where("user_id = ?", id).Find(&userTenors)
	if result.Error != nil {
		return userTenors, result.Error
	}

	return userTenors, nil
}
