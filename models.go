package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/har-sat/rssagg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

// this is just to not have camelcase in response and have custom json tags defined
func DatabaseUserToUser(usr database.User) User {
	return User{
		ID: usr.ID,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
		Name: usr.Name,
	}
}