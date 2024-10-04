package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"tickets/internal/application/interfaces"
	"tickets/internal/domain/entities"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Create(c context.Context, user entities.User) error {
	result := u.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *UserRepository) FindByEmail(c context.Context, email string) (entities.User, error) {
	var user entities.User
	result := u.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, entities.ErrNotFound
		}
		return user, result.Error
	}

	return user, nil
}

func (u *UserRepository) FindByUuid(c context.Context, uuid string) (entities.User, error) {
	var user entities.User
	result := u.db.Where("uuid = ?", uuid).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (u *UserRepository) Update(c context.Context, user entities.User) error {
	result := u.db.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *UserRepository) Delete(c context.Context, uuid string) error {
	result := u.db.Where("uuid = ?", uuid).Delete(&entities.User{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
