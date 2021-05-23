package services

import (
	"github.com/MaksymVakuliuk/golang-microservices/mvc/models"
	"github.com/MaksymVakuliuk/golang-microservices/mvc/utils"
)

func GetUser(userId uint64) (*models.User, *utils.ApplicationError) {
	return models.GetUser(userId)
}
