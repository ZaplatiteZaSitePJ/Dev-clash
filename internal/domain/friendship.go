package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type FriendshipStatus string

const (
	StatusPending  FriendshipStatus = "pending"
	StatusAccepted FriendshipStatus = "accepted"
	StatusRejected  FriendshipStatus = "rejected"
)

type Friendship struct {
	RequesterID uuid.UUID        `db:"requester_id"`
	AddresseeID uuid.UUID        `db:"addressee_id"`
	Status      FriendshipStatus `db:"status"`
	CreatedAt   time.Time        `db:"created_at"`
	UpdatedAt   time.Time        `db:"updated_at"`
}

func (f *Friendship) ValidateFriendship() error {
	if f.Status != StatusPending && f.Status != StatusAccepted && f.Status != StatusRejected {
		return errors.New("invalid friendship status")
	}
	return nil
} 