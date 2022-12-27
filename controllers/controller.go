package handler

import (
	"odoo-travel/services"

	"github.com/go-playground/validator/v10"
)

type handler struct {
	srv       services.Services
	validator *validator.Validate
}

func NewController(service services.Services, validator *validator.Validate) *handler {
	return &handler{
		srv:       service,
		validator: validator,
	}
}
