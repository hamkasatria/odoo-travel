package services

import (
	r "odoo-travel/repositories"
)

type Services interface {
}

type services struct {
	repo r.Repository
}

func NewService(repo r.Repository) Services {
	return &services{
		repo: repo,
	}
}
