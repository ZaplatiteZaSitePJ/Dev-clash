package user

import (
	"dev-clash/internal/domain"
)

type UserRepository interface {
	Save(*domain.User) (*domain.User, error)
	FindByID(id int) (*domain.User, error)
	FindAll() ([]*domain.User, error)
}