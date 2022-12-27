package services

import (
	"odoo-travel/models"
	r "odoo-travel/repositories"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type Services interface {
	GetListTravels(ctx *fiber.Ctx) (*[]models.Travel, error)
	GetTravelById(ctx *fiber.Ctx, id string) (*models.Travel, error)
	AddTravel(ctx *fiber.Ctx, travel *models.Travel) error
	EditTravel(ctx *fiber.Ctx, id string, travel *models.Travel) error
}

type services struct {
	repo r.Repository
}

func NewService(repo r.Repository) Services {
	return &services{
		repo: repo,
	}
}

func (s *services) GetListTravels(ctx *fiber.Ctx) (*[]models.Travel, error) {
	data, err := s.repo.GetListTravels(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *services) GetTravelById(ctx *fiber.Ctx, id string) (*models.Travel, error) {
	data, err := s.repo.GetTravelById(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *services) AddTravel(ctx *fiber.Ctx, travel *models.Travel) error {
	if err := s.repo.AddTravel(ctx, travel); err != nil {
		return err
	}

	return nil
}

func (s *services) EditTravel(ctx *fiber.Ctx, id string, travel *models.Travel) error {

	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: travel.Name},
				{Key: "contact", Value: travel.Contact},
			},
		},
	}

	if err := s.repo.EditTravel(ctx, id, update); err != nil {
		return err
	}

	return nil
}
