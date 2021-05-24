package services

import (
	"net/http"

	"github.com/MaksymVakuliuk/golang-microservices/mvc/models"
	"github.com/MaksymVakuliuk/golang-microservices/mvc/utils"
)

type itemService struct {
}

var (
	ItemService itemService
)

func (s *itemService) GetItem(itemId string) (*models.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
