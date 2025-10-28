package dto

type CreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SafetyUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}