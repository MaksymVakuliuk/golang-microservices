package services

import (
	"github.com/MaksymVakuliuk/golang-microservices/mvc/models"
	"github.com/MaksymVakuliuk/golang-microservices/mvc/utils"
)

type userService struct {
}

var (
	UserService userService
)

func (s *userService) GetUser(userId uint64) (*models.User, *utils.ApplicationError) {
	return models.UserDao.GetUser(userId)
}
