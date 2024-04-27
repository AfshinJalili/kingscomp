package entity

import (
	"time"
)

type Account struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	UserName  string    `json:"user_name"`
	JoinedAt  time.Time `json:"joined_at"`

	DisplayName string `json:"display_name"`
}

func (a Account) EntityID() ID {
	return NewID("account", a.ID)
}
