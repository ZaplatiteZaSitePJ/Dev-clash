package repositories

import (
	"database/sql"
	"dev-clash/internal/domain"
	"dev-clash/pkg/logger"
	"dev-clash/pkg/server_utils/app_errors"
	pg_err "dev-clash/pkg/server_utils/db_errors/postgres"

	"github.com/google/uuid"
	"github.com/lib/pq"
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
func (userRepo *UserRepository) Save(newUser *domain.User) (*domain.User, error) {
	var savedUser = &domain.User{}

	query := 
		`INSERT INTO users 
		(id, username, email, hashed_password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, username, email, rating, created_at, updated_at`
		
	err := userRepo.db.QueryRow(
		query, 
		newUser.ID, 
		newUser.Username, 
		newUser.Email, 
		newUser.HashedPassword, 
		newUser.CreatedAt, 
		newUser.UpdatedAt,
	).Scan(
		&savedUser.ID, 
		&savedUser.Username, 
		&savedUser.Email, 
		&savedUser.Rating, 
		&savedUser.CreatedAt, 
		&savedUser.UpdatedAt)

	if err != nil {
		logger.Error("db", err)
		if pg_err.IsUniqueViolation(err) {
			return nil, app_errors.AlreadyExists("user already exist", err)
		}
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}
	return savedUser, nil
}

// FIND USER BY ID IN DATABASE
func (userRepo *UserRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	findedUser := &domain.User{}
	
	userQuery := "SELECT id, username, email, rating FROM users WHERE id = $1"

	friendIDsQuery := `SELECT
			CASE
				WHEN requester_id = $1 THEN addressee_id
				ELSE requester_id
			END AS friends_id
		FROM friendship
		WHERE (requester_id = $1 OR addressee_id = $1) AND status=$2
		`

	err := userRepo.db.QueryRow(userQuery, id).Scan(&findedUser.ID, &findedUser.Username, &findedUser.Email, &findedUser.Rating)

	if err != nil {
		logger.Error("db", err)
		if err == sql.ErrNoRows {
			return nil, app_errors.NotFound("user not found", err)
		}
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}

	rows, err:= userRepo.db.Query(friendIDsQuery, id, domain.StatusAccepted)

	if err != nil {
		logger.Error("db", err)
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}

	defer rows.Close()

	findedUser.FriendIDs = []uuid.UUID{}
	for rows.Next() {
		var friendID uuid.UUID

		err := rows.Scan(&friendID)

		if err != nil {
			logger.Error("db", err)
			return nil, app_errors.Internal("server unavailable now. Try again later", err)
		}

		findedUser.FriendIDs = append(findedUser.FriendIDs, friendID)
	}

	if err := rows.Err(); err != nil {
		logger.Error("db", err)
        return nil, app_errors.Internal("server unavailable now. Try again later", err)
    }

	return findedUser, nil
}

func (userRepo *UserRepository) FindBySeveralIDs(ids []uuid.UUID) ([]*domain.User, error) {
	if len(ids) == 0 {
        return []*domain.User{}, nil
    }

	findedUsers := make([]*domain.User, 0, len(ids))

	query :=
		`
		SELECT id, username, email, rating FROM users
		WHERE id = ANY($1)
		`

	rows, err:= userRepo.db.Query(query, pq.Array(ids))

	if err != nil {
		logger.Error("db", err)
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}

	defer rows.Close()

	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Rating)
		
		if err != nil {
			logger.Error("db", err)
			return nil, app_errors.Internal("server unavailable now. Try again later", err)
		}

		findedUsers = append(findedUsers, user)
	}

	return findedUsers, nil
}

func (userRepo *UserRepository) FindAll() ([]*domain.User, error) {
	
	findedUsers := []*domain.User{}

	query := "SELECT id, username, email, rating FROM users"

	rows, err := userRepo.db.Query(query)
	if err != nil {
		logger.Error("db", err)
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}

	defer rows.Close()

	for rows.Next() {
		user := &domain.User{}
		rows.Scan(&user.ID, &user.Username, &user.Email, &user.Rating)
		findedUsers = append(findedUsers, user)
	}

	if err := rows.Err(); err != nil {
		logger.Error("db", err)
        return nil, app_errors.Internal("server unavailable now. Try again later", err)
    }

	return findedUsers, nil
}

func (userRepo *UserRepository) DeleteByID(id int) error {
	return nil
}