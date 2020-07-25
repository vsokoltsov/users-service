package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User defines a user representation
type User struct {
	BaseModel
	ID        int       `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// SetPassword hash given password and save it to the user
func (u *User) SetPassword(password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = string(hashedPassword)
}
