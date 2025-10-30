package dto

import (
	"dev-clash/internal/domain"
)

type CreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SafetyUser struct {
	Username         string   `json:"username"`
	Email            string   `json:"email"`
	Skills           []string `json:"skills,omitempty"`
	ParticipantTimes int      `json:"participant_times"`
	PrizeTimes       int      `json:"prize_times"`
	ModeratorTimes   int      `json:"moderator_times"`
	Status           string   `json:"status,omitempty"`
	Description      string   `json:"description,omitempty"`
}

func SafetyUserFromModel(user *domain.User) *SafetyUser {
	return &SafetyUser{
		Username: user.Username, 
		Email: user.Email, 
		Description: user.Description,
		Status: user.Status,
		ModeratorTimes: user.ModeratorTimes,
		ParticipantTimes: user.ParticipantTimes,
		PrizeTimes: user.PrizeTimes,
		Skills: user.Skills,
	}
}