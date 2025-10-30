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
func (userRepo *UserRepository) FindByID(id int) (*user.User, error) {
	query := 
		`SELECT u.id, u.username, u.email, u.description, u.moderators_times, u.prizes_times, u.participant_times, u.status, s.title
		FROM users u
		LEFT JOIN users_skills us ON us.user_id = u.id
		LEFT JOIN skills s ON us.skill_id = s.id
		WHERE u.id = $1
	`

	var findedUser = &user.User{}

	rows, err := userRepo.db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("db error: %w", err)
	}

	defer rows.Close()

	firstItter := true 
	for rows.Next() {
		
		var username, email string
		var description, status, skill sql.NullString
		var id, moderators_times, prizes_times, participant_times int

		rows.Scan(&id, &username, &email, &description, &moderators_times, &prizes_times, &participant_times, &status, &skill)

		if firstItter {
			findedUser.ID = id
			findedUser.PrizeTimes = prizes_times
			findedUser.ParticipantTimes = participant_times
			findedUser.Description = description
			findedUser.Status = status
			findedUser.ModeratorTimes = moderators_times
			findedUser.Email = email
			findedUser.Username = username
			firstItter = false
		}
		findedUser.Skills = append(findedUser.Skills, skill)
	} 

	if firstItter {
		return nil, fmt.Errorf("user not found")
	}

	return findedUser, nil
}