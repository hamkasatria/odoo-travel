package services

import (
	"odoo-travel/models"
	r "odoo-travel/repositories"

	"github.com/gofiber/fiber/v2"
)

type Services interface {
	GetListTravels(ctx *fiber.Ctx) (*[]models.Travel, error)
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
