package transports

import (
	"odoo-travel/models"
)

type GeneralResponse struct {
	Message string `json:"message"`
}

type GetListTravels struct {
	Count       int             `json:"count"`
	ListTravels []models.Travel `json:"list_travels"`
}

type ResponseError struct {
	Message string
	Status  int
}
