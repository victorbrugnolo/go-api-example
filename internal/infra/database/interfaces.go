package database

import "github.com/victorbrugnolo/go-api-example/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
