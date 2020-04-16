package forms

import (
	"github.com/go-playground/validator/v10"
	"github.com/vsokoltsov/users-service/app/models"
	"github.com/vsokoltsov/users-service/app/utils"
)

// UserForm represents user
// parameters for create or update operation
type UserForm struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"email,required"`
	Password  string `validate:"email,required"`
}

func (uf *UserForm) validate() error {
	var (
		validator = validator.New()
	)
	err := validator.Struct(uf)
	return err
}

// Submit perform saving of users data
func (uf *UserForm) Submit() (*models.User, error) {
	var user = models.User{}
	validationError := uf.validate()

	if validationError != nil {
		return &user, validationError
	}
	user.SetPassword(uf.Password)
	tx := utils.DB.MustBegin()
	tx.QueryRowx(
		"insert into users(first_name, last_name, email, password) values ($1, $2, $3, $4) returning *",
		uf.FirstName,
		uf.LastName,
		uf.Email,
		user.Password,
	).StructScan(&user)
	err := tx.Commit()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
