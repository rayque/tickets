package repositories

import (
	"context"
	"tickets/internal/application/interfaces"
	"tickets/internal/domain/entities"
)

type UserRepository struct{}

func NewUserRepository() interfaces.UserRepository {
	return &UserRepository{}
}
func (u *UserRepository) Create(c context.Context, user entities.User) error {
	return nil
}

func (u *UserRepository) FindByUuid(c context.Context, uuid string) (entities.User, error) {
	return entities.User{}, nil
}

func (u *UserRepository) Update(c context.Context, user entities.User) error {
	return nil
}

func (u *UserRepository) Delete(c context.Context, uuid string) error {
	return nil
}
