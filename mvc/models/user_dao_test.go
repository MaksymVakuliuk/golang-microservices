package models

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)
	assert.Nil(t, user, "We were not expecting a user with id 0")
	assert.NotNil(t, err, "We were expecting an error when user id is 0")
	assert.EqualValues(t, "user 0 was not found", err.Message)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err, "We were not expecting error")
	assert.NotNil(t, user, "We were expecting no nil user")
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Ivan", user.FirstName)
	assert.EqualValues(t, "Ivanov", user.SecondName)
	assert.EqualValues(t, "iviv@gmail.com", user.Email)

}
