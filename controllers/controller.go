package controllers

import (
	"odoo-travel/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type controller struct {
	srv       services.Services
	validator *validator.Validate
}

func NewController(service services.Services, validator *validator.Validate) *controller {
	return &controller{
		srv:       service,
		validator: validator,
	}
}

func (c *controller) GetListTravel(ctx *fiber.Ctx) error {
	// fb.Header("Content-Type", "application/json")
	result, err := c.srv.GetListTravels(ctx)

	if err != nil {
		ctx.JSON(err)
		return nil
	}

	ctx.JSON(result)
	return nil
}
