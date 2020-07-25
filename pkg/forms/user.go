package forms

import (
	"github.com/vsokoltsov/users-service/pkg/models"
	"github.com/vsokoltsov/users-service/pkg/utils"
)

// UserForm represents user
// parameters for create or update operation
type UserForm struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `validate:"required"`
}

// validate call default form validation method
func (uf *UserForm) validate() map[string][]string {
	return DefaultFormValidation(uf)
}

// Submit perform saving of users data
func (uf *UserForm) Submit() (models.DBModel, *map[string][]string) {
	var (
		user            = models.User{}
		validationError = uf.validate()
	)

	// If validation has failed
	if len(validationError) > 0 {
		return &user, &validationError
	}
	user.SetPassword(uf.Password)

	// Add new user to database
	tx := utils.DB.MustBegin()
	tx.QueryRowx(
		"insert into users(first_name, last_name, email, password) values ($1, $2, $3, $4) returning *",
		uf.FirstName,
		uf.LastName,
		uf.Email,
		user.Password,
	).StructScan(&user)

	// If adding failed - return errors
	err := tx.Commit()
	if err != nil {
		return nil, &map[string][]string{
			"user": []string{err.Error()},
		}
	}
	return user, nil
}
