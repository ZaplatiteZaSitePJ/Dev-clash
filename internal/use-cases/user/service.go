package user

import (
	"dev-clash/internal/domain"
	"dev-clash/internal/dto"
	"dev-clash/pkg/crypt_password"
	"dev-clash/pkg/logger"
	custom_errors "dev-clash/pkg/server_utils/errors"
	"fmt"
	"strings"
)

type UserService struct {
	repo UserRepository
}

func New(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}


func (u *UserService) CreateUser(input *dto.CreateUser) (*domain.User, error){
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

	newUser := domain.User{
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

func (u *UserService) FindUserByID(userID int) (*domain.User, error){
	logger.Info("Trying to find user: ", userID)

	findedUser, err := u.repo.FindByID(userID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			wErr := custom_errors.New(err, 404)
			wErr.AddResponseData("User not found")
			wErr.AddLogData(fmt.Sprintf("user with id=%d not found", userID))
			return nil, wErr
		}

		wErr := custom_errors.New(err, 500)
		wErr.AddResponseData("Internal server error")
		wErr.AddLogData(err.Error())
		return nil, wErr
	}
	return findedUser, nil
}
