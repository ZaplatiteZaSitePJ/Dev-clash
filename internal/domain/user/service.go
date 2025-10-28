package user

import (
	"dev-clash/internal/dto"
	"dev-clash/pkg/crypt_password"
	"dev-clash/pkg/logger"
	custom_errors "dev-clash/pkg/server_utils/errors"
	"fmt"
)

type UserService struct {
	repo UserRepository
}

func New(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}


func (u *UserService) CreateUser(input *dto.CreateUser) (*User, error){
	logger.Info("Creating new user: ", input)

	var hashed_password string

	// Validation password
	if err := crypt_password.ValidatePassword(input.Password); err != nil {
		wError := custom_errors.New(err, 422)
		wError.AddLogData(fmt.Sprintf("week password: %+v", input))
		wError.AddResponseData("password too week")
		return nil, wError 
	}

	// Hashing password
	hashed_password, err := crypt_password.EncryptPassword(input.Password) 
	if err != nil {
		wError := custom_errors.New(err, 500)
		wError.AddLogData(fmt.Sprintf("failed to hashing password: %+v", input.Password))
		wError.AddResponseData("Sorry, we gave some troubles, try again later")
		return nil, wError 
	}

	newUser := User{
		Email: input.Email,
		Username: input.Username,
		HashedPassword: hashed_password,
	}

	// Validation user
	if err := newUser.Validate(); err != nil {
		wError := custom_errors.New(err, 422)
		wError.AddLogData(fmt.Sprintf("Incorrect user data: %+v", input))
		wError.AddResponseData("Invalid data, correct user form")
		return nil, wError 
	}

	return u.repo.Save(&newUser)
}