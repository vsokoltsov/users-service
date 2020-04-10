package forms

// UserForm represents user
// parameters for create or update operation
type UserForm struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string
}
