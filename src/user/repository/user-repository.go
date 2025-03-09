package repository

import (
	"errors"
	"test-kp-golang/src/user/entity"
)

type UserRepository struct {
	users  map[int]entity.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  make(map[int]entity.User),
		nextID: 1,
	}
}

func (r *UserRepository) CreateUser(user entity.User) (entity.User, error) {
	user.ID = r.nextID
	r.users[user.ID] = user
	r.nextID++
	return user, nil
}

func (r *UserRepository) GetUserByID(id int) (entity.User, error) {
	user, exists := r.users[id]
	if !exists {
		return entity.User{}, errors.New("user not found")
	}
	return user, nil
}
