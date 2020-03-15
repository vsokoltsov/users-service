package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var buffer bytes.Buffer
	buffer.WriteString(`{Response: "success", Message: "users service"}`)
	json.NewEncoder(w).Encode(buffer.String())
	w.WriteHeader(http.StatusOK)
	return
}
