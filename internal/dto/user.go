package dto

import "database/sql"

func NullStringToValid(ns sql.NullString) string{
	if !ns.Valid {
		return ""
	}

	return  ns.String
}

func NullStringSliceToValid(ns []sql.NullString) []string{
	strs := make([]string, 0, 5)
	
	for _, n := range(ns) {
		if !n.Valid {
			break 
		} else {
			strs = append(strs, n.String)
		}	
	}

	return strs
}


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
	Status           string   `json:"status"`
	Description      string   `json:"description"`
}
