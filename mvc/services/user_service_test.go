package services

import (
	"net/http"
	"testing"

	"github.com/MaksymVakuliuk/golang-microservices/mvc/models"
	"github.com/MaksymVakuliuk/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock     usersDaoMock
	getUserFunction func(userId uint64) (*models.User, *utils.ApplicationError)
)

func init() {
	models.UserDao = &userDaoMock
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userId uint64) (*models.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNoFoundUserInDatabase(t *testing.T) {
	getUserFunction = func(userId uint64) (*models.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "user 0 was not found",
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	user, err := userDaoMock.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
	assert.EqualValues(t, "not_found", err.Code)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId uint64) (*models.User, *utils.ApplicationError) {
		return &models.User{
			Id:         123,
			FirstName:  "Ivan",
			SecondName: "Ivanov",
			Email:      "iviv@gmail.com",
		}, nil
	}
	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Ivan", user.FirstName)
	assert.EqualValues(t, "Ivanov", user.SecondName)
	assert.EqualValues(t, "iviv@gmail.com", user.Email)

}
