package usecase

import "test-kp-golang/src/domain/user/entity"

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUserByID(id int) (entity.User, error)
}

type UserUsecase struct {
	userRepo UserRepository
}

func NewUserUsecase(userRepo UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) CreateUser(user entity.User) (entity.User, error) {
	return u.userRepo.CreateUser(user)
}

func (u *UserUsecase) GetUserByID(id int) (entity.User, error) {
	return u.userRepo.GetUserByID(id)
}
