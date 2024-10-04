package interfaces

import (
	"context"
	"tickets/internal/domain/entities"
)

type UserUseCase interface {
	Create(c context.Context, user entities.User) error
	FindByUuid(c context.Context, uuid string) (entities.User, error)
	Update(c context.Context, user entities.User) error
	Delete(c context.Context, uuid string) error
}
