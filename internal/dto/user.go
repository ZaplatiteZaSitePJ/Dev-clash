package dto

import (
	"dev-clash/internal/domain"

	"github.com/shopspring/decimal"
)

type CreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PublicUser struct {
	Username         string   `json:"username"`
	Email            string   `json:"email"`
	Rating         decimal.Decimal `db:"rating" json:"rating"`
}

type PublicUserWithFriends struct {
	Username         string   `json:"username"`
	Email            string   `json:"email"`
	Rating         decimal.Decimal `db:"rating" json:"rating"`
	Friends 		[]*PublicUser `json:"friends"`
}

func PublicUserWithFriendsFromModel(user *domain.User, friends []*domain.User) *PublicUserWithFriends {
	convertedFriends := SeveralUsersToPublic(friends)
	return &PublicUserWithFriends{
		Username: user.Username, 
		Email: user.Email,
		Rating: user.Rating, 
		Friends: convertedFriends,
	}
}

func PublicUserFromModel(user *domain.User) *PublicUser {
	return &PublicUser{
		Username: user.Username, 
		Email: user.Email,
		Rating: user.Rating, 
	}
}

func SeveralUsersToPublic(users []*domain.User) []*PublicUser {
    publicUsers := make([]*PublicUser, 0, len(users))
    for _, u := range users {
        publicUsers = append(publicUsers, PublicUserFromModel(u))
    }
    return publicUsers
}