package models

import (
	"time"

	_ "github.com/lib/pq"
)

// User defines a user representation
type User struct {
	ID         int       `db:"id"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	Email      string    `db:"email"`
	CreatedAt  time.Time `db:"created_at"`
	UpdateddAt time.Time `db:"updated_at"`
}
