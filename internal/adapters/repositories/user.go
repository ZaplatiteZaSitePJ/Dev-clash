package repositories

import (
	"database/sql"
	"dev-clash/internal/domain/user"
	"dev-clash/pkg/logger"
	custom_errors "dev-clash/pkg/server_utils/errors"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// SAVE USER IN DATABASE
func (userRepo *UserRepository) Save(newUser *user.User) (*user.User, error) {
	logger.Info("Trying to push new user in DB", newUser)
	query := `INSERT INTO users (username, email, hashed_password) VALUES ($1, $2, $3) RETURNING id`
	if err := userRepo.db.QueryRow(query, newUser.Username, newUser.Email, newUser.HashedPassword).Scan(&newUser.ID); err != nil {
		wError := custom_errors.New(err, 500)
		wError.AddLogData(fmt.Sprintf("failed to push user in db: %+v", newUser))
		wError.AddResponseData("Sorry, we gave some troubles, try again later")
		return nil, wError 
	}

	return newUser, nil
}

// FIND USER BY ID IN DATABASE
func (userRepo *UserRepository) FindByID(id int) (*user.User, bool, error) {
	query := `SELECT * FROM users WHERE id = $1 RETURNING id, username, email`
	var findedUser user.User

	if err := userRepo.db.QueryRow(query, id).Scan(&findedUser.ID, &findedUser.Username, &findedUser.Email); err != nil {
		return &user.User{}, false, err
	}

	return &findedUser, true, nil
}