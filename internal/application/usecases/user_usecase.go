package usecases

import (
	"context"
	"tickets/internal/application/interfaces"
	"tickets/internal/domain/entities"
)

type UserUseCase struct {
	UserRepository interfaces.UserRepository
}

func NewUserUseCase(userRepository interfaces.UserRepository) interfaces.UserUseCase {
	return &UserUseCase{UserRepository: userRepository}
}

func (u *UserUseCase) Create(c context.Context, user entities.User) error {
	return u.UserRepository.Create(c, user)
}

func (u *UserUseCase) FindByUuid(c context.Context, uuid string) (entities.User, error) {
	return u.UserRepository.FindByUuid(c, uuid)
}

func (u *UserUseCase) Update(c context.Context, user entities.User) error {
	return u.UserRepository.Update(c, user)
}

func (u *UserUseCase) Delete(c context.Context, uuid string) error {
	return u.UserRepository.Delete(c, uuid)
}
