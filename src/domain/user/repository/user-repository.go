package repository

import (
	"test-kp-golang/src/domain/user/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user entity.User) (entity.User, error) {
	r.db.Create(&user)
	return user, nil
}

func (r *UserRepository) GetUserByID(id int) (entity.User, error) {
	var user entity.User

	result := r.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User

	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
