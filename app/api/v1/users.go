package v1

import (
	"encoding/json"
	"net/http"

	"github.com/vsokoltsov/users-service/app/models"
	"github.com/vsokoltsov/users-service/app/serializers"
	"github.com/vsokoltsov/users-service/app/utils"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var serializedUsers = []serializers.UserSerializer{}
	var users = []models.User{}
	utils.DB.Select(&users, "select * from users")
	for _, user := range users {
		serializedUsers = append(serializedUsers, serializers.UserSerializer{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		})
	}
	json.NewEncoder(w).Encode(serializedUsers)
}
