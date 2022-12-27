package controllers

import (
	"odoo-travel/models"
	"odoo-travel/services"

	transports "odoo-travel/transports"

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
		transports.JsonResponse(ctx, fiber.StatusUnprocessableEntity, err.Error())
		return nil
	}

	trans := &transports.GetListTravels{
		Count: len(*result),
		List:  *result,
	}

	transports.JsonResponse(ctx, fiber.StatusOK, trans)

	return nil
}

func (c *controller) GetTravelById(ctx *fiber.Ctx) error {
	id := ctx.Params("ObjectId")
	result, err := c.srv.GetTravelById(ctx, id)
	if err != nil {
		transports.JsonResponse(ctx, fiber.StatusUnprocessableEntity, err.Error())
		return nil
	}

	if result.ID == "" {
		transports.JsonResponse(ctx, fiber.StatusNotFound, fiber.ErrNotFound)
		return nil
	}

	transports.JsonResponse(ctx, fiber.StatusOK, result)
	return nil
}

func (c *controller) AddTravel(ctx *fiber.Ctx) error {

	travel := new(transports.InsertTravel)

	if err := ctx.BodyParser(travel); err != nil {
		transports.JsonResponse(ctx, fiber.StatusBadRequest, err.Error())
		return err
	}

	if err := c.validator.Struct(travel); err != nil {
		transports.JsonResponse(ctx, fiber.StatusBadRequest, err.Error())
		return nil
	}

	if err := c.srv.AddTravel(ctx, &models.Travel{
		Name:    travel.Name,
		Contact: travel.Contact,
	}); err != nil {
		transports.JsonResponse(ctx, fiber.StatusUnprocessableEntity, err.Error())
		return nil
	}

	transports.JsonResponse(ctx, fiber.StatusOK, nil)
	return nil
}
