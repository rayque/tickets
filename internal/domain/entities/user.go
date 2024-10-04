package entities

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

var ErrUserAlreadyExists = errors.New("user already exists")
var ErrNotFound = errors.New("not found")

type User struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	Uuid      string
	Email     string `gorm:"index:idx_name,unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	gorm.Model
}
