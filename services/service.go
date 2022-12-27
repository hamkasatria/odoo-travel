package services

import (
	r "odoo-travel/repositories"
	"odoo-travel/transports"

	"github.com/gofiber/fiber/v2"
)

type Services interface {
	GetListTravels(ctx *fiber.Ctx) (*transports.GetListTravels, *transports.ResponseError)
}

type services struct {
	repo r.Repository
}

func NewService(repo r.Repository) Services {
	return &services{
		repo: repo,
	}
}

func (s *services) GetListTravels(ctx *fiber.Ctx) (*transports.GetListTravels, *transports.ResponseError) {
	data, err := s.repo.GetListTravels(ctx)
	if err != nil {
		return nil, nil
	}
	trans := &transports.GetListTravels{
		Count:       len(*data),
		ListTravels: *data,
	}
	return trans, nil
}
