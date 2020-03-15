package v1

import (
	"encoding/json"
	"github.com/vsokoltsov/users-service/app/serializers"
	"net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := serializers.GetUserSerializer()
	json.NewEncoder(w).Encode(user)
}
