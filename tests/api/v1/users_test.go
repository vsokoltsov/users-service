package apitests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/vsokoltsov/users-service/pkg/models"
	"github.com/vsokoltsov/users-service/pkg/utils"
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
		"password":   "password",
	}
	tx := utils.DB.MustBegin()
	tx.QueryRowx(
		"insert into users(first_name, last_name, email, password) values ($1, $2, $3, $4) returning *",
		udata["first_name"],
		udata["last_name"],
		udata["email"],
		udata["password"],
	).StructScan(&u)
	err := tx.Commit()
	if err != nil {
		t.Error("Error fo saving to database", err)
		log.Fatalln(err)
	}
	var receivedUsers []models.User

	response := tests.MakeRequest("GET", "/api/v1/users", nil)
	json.Unmarshal(response.Body.Bytes(), &receivedUsers)

	if response.Code != http.StatusOK {
		t.Error("Response status is not success")
	}

	if len(receivedUsers) != 1 {
		t.Error("Users list has not been received")
	}
	if u.ID != receivedUsers[0].ID {
		t.Error("Users do not match")
	}
}

// Test POST /api/v1/users - success user creation
func TestSuccessUserCreation(t *testing.T) {
	var (
		params = []byte(`{
		"email": "test@mail.com", 
		"first_name": "test", 
		"last_name": "test", 
		"password": "password"
	}`)
		usersCount            int
		usersCountAfterCreate int
	)
	err := utils.DB.Get(&usersCount, "select count(*) from users")
	if err != nil {
		t.Error("Cannot get a user's count: ", err)
	}
	response := tests.MakeRequest("POST", "/api/v1/users", bytes.NewBuffer(params))

	if response.Code != http.StatusCreated {
		t.Error("Response status is not success")
	}
	createErr := utils.DB.Get(&usersCountAfterCreate, "select count(*) from users")
	if createErr != nil {
		t.Error("Cannot get a user's count after create: ", createErr)
	}
	t.Log(usersCount, usersCountAfterCreate)
	if usersCountAfterCreate != usersCount+1 {
		t.Error("POST /api/v1/users failed: Number of users does not increased")
	}
}

// Test POST /api/v1/users - failed user creation
func TestFailedUserCreation(t *testing.T) {
	var (
		params                = []byte(`{}`)
		usersCount            int
		usersCountAfterCreate int
	)
	err := utils.DB.Get(&usersCount, "select count(*) from users")
	if err != nil {
		t.Error("Cannot get a user's count: ", err)
	}
	response := tests.MakeRequest("POST", "/api/v1/users", bytes.NewBuffer(params))

	if response.Code != http.StatusBadRequest {
		t.Error("Response status is not 'Bad Request'")
	}

	createErr := utils.DB.Get(&usersCountAfterCreate, "select count(*) from users")
	if createErr != nil {
		t.Error("Cannot get a user's count after create: ", createErr)
	}
	t.Log(usersCount, usersCountAfterCreate)
	if usersCountAfterCreate != usersCount {
		t.Error("POST /api/v1/users failed: Number of users has changed")
	}
}
