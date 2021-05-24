package models

import (
	"fmt"
	"net/http"

	"github.com/MaksymVakuliuk/golang-microservices/mvc/utils"
)

var (
	users = map[uint64]*User{
		123: {Id: 123, FirstName: "Ivan", SecondName: "Ivanov", Email: "iviv@gmail.com"},
	}
)

func GetUser(userId uint64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
