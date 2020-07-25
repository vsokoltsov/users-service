package forms

import (
	"reflect"

	"github.com/go-playground/validator/v10"
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

func (uf *UserForm) validate() map[string][]string {
	var (
		valid   = validator.New()
		errsMap = make(map[string][]string)
	)
	fieldValues, fieldTags := getFieldsWithValues(uf)
	for key, value := range fieldValues {
		tag := fieldTags[key]
		ferr := valid.Var(value, tag.(string))
		if ferr != nil {
			var errStrings []string
			errsData := ferr.(validator.ValidationErrors)
			for _, errItem := range errsData {
				errStrings = append(errStrings, errItem.Tag())
			}
			errsMap[key] = errStrings
		}
	}
	return errsMap
}

func getFieldsWithValues(uf *UserForm) (map[string]interface{}, map[string]interface{}) {
	var (
		fieldValue = make(map[string]interface{})
		fieldTag   = make(map[string]interface{})
	)
	rfields := reflect.TypeOf(*uf)
	rvalues := reflect.ValueOf(uf).Elem()
	for i := 0; i < rfields.NumField(); i++ {
		field := rfields.Field(i)
		value := rvalues.Field(i)
		tag := field.Tag.Get("validate")
		fieldValue[field.Name] = value.Interface()
		fieldTag[field.Name] = tag
	}
	return fieldValue, fieldTag
}

// Submit perform saving of users data
func (uf *UserForm) Submit() (*models.User, *map[string][]string) {
	var user = models.User{}
	validationError := uf.validate()

	if len(validationError) > 0 {
		return &user, &validationError
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
		return nil, &map[string][]string{
			"user": []string{err.Error()},
		}
	}
	return &user, nil
}
