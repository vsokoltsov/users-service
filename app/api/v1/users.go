package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vsokoltsov/users-service/app/forms"
	"github.com/vsokoltsov/users-service/app/utils"

	"github.com/vsokoltsov/users-service/app/models"
	"github.com/vsokoltsov/users-service/app/serializers"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var serializedUsers = []serializers.UserSerializer{}
	var users []models.User
	utils.DB.Select(&users, "SELECT * FROM users")
	for _, user := range users {
		serializedUsers = append(serializedUsers, serializers.UserSerializer{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		})
	}
	json.NewEncoder(w).Encode(serializedUsers)
}

func createUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var form forms.UserForm
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&form)

	tx := utils.DB.MustBegin()
	tx.MustExec(
		"insert into users(first_name, last_name, email) values ($1, $2, $3)",
		form.FirstName,
		form.LastName,
		form.Email,
	)
	err := tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
}
