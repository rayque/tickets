package usecases

import (
	"context"
	"errors"
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
	resultUser, err := u.UserRepository.FindByEmail(c, user.Email)
	if err != nil && !errors.Is(err, entities.ErrNotFound) {
		return err
	}

	if resultUser.Email != "" {
		return entities.ErrUserAlreadyExists
	}

	return u.UserRepository.Create(c, entities.User{
		Uuid:  user.Uuid,
		Email: user.Email,
	})
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
