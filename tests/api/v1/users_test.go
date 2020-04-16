package api_tests

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/vsokoltsov/users-service/app/models"
	"github.com/vsokoltsov/users-service/app/utils"
	"github.com/vsokoltsov/users-service/tests"
)

func TestMain(m *testing.M) {
	tests.TestMain(m)
}

// Test GET /api/v1/users path
func TestApiUsersGetRoute(t *testing.T) {
	var u models.User

	var udata = map[string]string{
		"first_name": "test",
		"last_name":  "test",
		"email":      "test@gmail.com",
	}
	tx := utils.DB.MustBegin()
	tx.QueryRowx(
		"insert into users(first_name, last_name, email) values ($1, $2, $3) returning id, first_name, last_name, email",
		udata["first_name"],
		udata["last_name"],
		udata["email"],
	).StructScan(&u)
	err := tx.Commit()
	if err != nil {
		t.Error("Error fo saving to database", err)
		log.Fatalln(err)
	}
	var receivedUsers []models.User

	response := tests.MakeRequest("GET", "/api/v1/users")
	json.Unmarshal(response.Body.Bytes(), &receivedUsers)

	if len(receivedUsers) != 1 {
		t.Error("Users list has not been received")
	}
	if u.ID != receivedUsers[0].ID {
		t.Error("Users do not match")
	}
}
