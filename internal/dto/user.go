package dto

import "database/sql"

func NullStringToValid(ns sql.NullString) string{
	if !ns.Valid {
		return ""
	}

	return  ns.String
}

type CreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SafetyUser struct {
	Username         string   `json:"username"`
	Email            string   `json:"email"`
	Skills           []string `json:"skills"`
	ParticipantTimes int      `json:"participant_times"`
	PrizeTimes       int      `json:"prize_times"`
	ModeratorTimes   int      `json:"moderator_times"`
	Status           string   `json:"status"`
	Description      string   `json:"description"`
}
