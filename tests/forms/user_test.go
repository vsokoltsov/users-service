package forms_test

import (
	"log"
	"testing"

	"github.com/vsokoltsov/users-service/app/forms"
	"github.com/vsokoltsov/users-service/app/utils"
	"github.com/vsokoltsov/users-service/tests"
)

func TestMain(m *testing.M) {
	tests.TestMain(m)
}

func getUsersCount() int {
	var usersCount int
	rows := utils.DB.QueryRow("select count(*) from users")
	err := rows.Scan(&usersCount)
	if err != nil {
		log.Fatal(err)
	}
	return usersCount
}

func TestSuccessUserCreation(t *testing.T) {

	var data = map[string]string{
		"email":    "test@gmail.com",
		"password": "awdawdwad",
	}
	var form = forms.UserForm{
		Email:    data["email"],
		Password: data["password"],
	}
	usersCount := getUsersCount()
	_, err := form.Submit()
	if err != nil {
		t.Error("Error appeared after creating user")
	}
	createdUsersCount := getUsersCount()
	if createdUsersCount != usersCount+1 {
		t.Error("User was not saved")
	}
}

func TestFailedUserCreation(t *testing.T) {
	var data = map[string]string{}
	var form = forms.UserForm{
		Email:    data["email"],
		Password: data["password"],
	}
	usersCount := getUsersCount()
	_, err := form.Submit()
	if err == nil {
		t.Error("Error was not appeared when it should")
	}
	createdUsersCount := getUsersCount()
	if createdUsersCount != usersCount {
		t.Error("Number of users dos not match after failed saving")
	}
}
