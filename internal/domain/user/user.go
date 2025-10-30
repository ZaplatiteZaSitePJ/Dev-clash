package user

import (
	"database/sql"
	"errors"
	"regexp"
	"strings"
	"unicode/utf8"
)

type User struct {
	ID             int		`json:"id"`
	Username       string	`json:"username"`
	Email          string	`json:"email"`
	HashedPassword string	`json:"hashed_password,omitempty"`
	ParticipantTimes int	`json:"participant_times,omitempty"`
	PrizeTimes       int	`json:"prize_times,omitempty"`
	ModeratorTimes   int	`json:"moderator_times,omitempty"`
	Status           sql.NullString `json:"status,omitempty"`
	Description      sql.NullString	`json:"description,omitempty"`
	Skills 	[]sql.NullString `json:"skills,omitempty"`
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func (u *User) Validate() error {
	trimedUsername := strings.TrimSpace(u.Username)

	if trimedUsername == "" || utf8.RuneCountInString(trimedUsername) < 3 {
		return errors.New("username cannot be empty or less 3 letter")
	}

	if !emailRegex.MatchString(u.Email) {
		return errors.New("invalid email format")
	}

	return nil
}

