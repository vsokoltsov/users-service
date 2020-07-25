package v1

import (
	"encoding/json"
	"net/http"

	"github.com/vsokoltsov/users-service/pkg/forms"
	"github.com/vsokoltsov/users-service/pkg/models"
	"github.com/vsokoltsov/users-service/pkg/utils"

	"github.com/vsokoltsov/users-service/pkg/serializers"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var serializedUsers = []serializers.UserSerializer{}
	var users []models.User
	utils.DB.Select(&users, "SELECT * FROM users")
	for _, user := range users {
		serializedUsers = append(serializedUsers, serializers.GetUserSerializer(user))
	}
	json.NewEncoder(w).Encode(serializedUsers)
}

func createUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var form forms.UserForm
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&form)

	user, err := form.Submit()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusCreated)
		serializer := serializers.GetUserSerializer(user.(models.User))
		json.NewEncoder(w).Encode(serializer)
	}
}
