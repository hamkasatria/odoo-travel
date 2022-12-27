package controllers

import (
	"odoo-travel/services"

	ress "odoo-travel/transports"

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
	result, err := c.srv.GetListTravels(ctx)

	if err != nil {
		ress.JsonResponse(ctx, fiber.StatusUnprocessableEntity, err.Error())
		return nil
	}

	trans := &ress.GetListTravels{
		Count: len(*result),
		List:  *result,
	}

	ress.JsonResponse(ctx, fiber.StatusOK, trans)

	return nil
}

func (c *controller) GetTravelById(ctx *fiber.Ctx) error {
	id := ctx.Params("ObjectId")
	result, err := c.srv.GetTravelById(ctx, id)
	if err != nil {
		ress.JsonResponse(ctx, fiber.StatusUnprocessableEntity, err.Error())
		return nil
	}

	if result.ID == "" {
		ress.JsonResponse(ctx, fiber.StatusNotFound, fiber.ErrNotFound)
		return nil
	}

	ress.JsonResponse(ctx, fiber.StatusOK, result)
	return nil
}
