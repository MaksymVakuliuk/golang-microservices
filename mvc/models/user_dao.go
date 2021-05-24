package models

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MaksymVakuliuk/golang-microservices/mvc/utils"
)

var (
	users = map[uint64]*User{
		123: {Id: 123, FirstName: "Ivan", SecondName: "Ivanov", Email: "iviv@gmail.com"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(uint64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (d *userDao) GetUser(userId uint64) (*User, *utils.ApplicationError) {
	log.Println("We're accessing to database")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
