package handlers

import "dev-clash/internal/domain/user"

type Handlers struct {
	User *user.UserService
}

func New(user *user.UserService) *Handlers {
	return &Handlers{
		User: user,
	}
}