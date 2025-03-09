package usecase

import "test-kp-golang/src/user/entity"

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUserByID(id int) (entity.User, error)
}

type UserUsecase struct {
	UserRepository UserRepository
}

func NewUserUsecase(userRepository UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
	}
}

func (u *UserUsecase) CreateUser(user entity.User) (entity.User, error) {
	return u.UserRepository.CreateUser(user)
}

func (u *UserUsecase) GetUserByID(id int) (entity.User, error) {
	return u.UserRepository.GetUserByID(id)
}
