package models

import "time"

// User defines a user representation
type User struct {
	ID         int        `db:"id"`
	FirstName  string     `db:"first_name"`
	LastName   string     `db:"last_name"`
	Email      string     `db:"email"`
	CreatedAt  time.Timer `db:"created_at"`
	UpdateddAt time.Timer `db:"updated_at"`
}
