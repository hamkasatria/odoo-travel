package server

import (
	con "odoo-travel/controllers"
	repo "odoo-travel/repositories"
	srv "odoo-travel/services"

	"github.com/gofiber/fiber/v2"
)

func (s *ApiServer) registerRouters() {

	repo := repo.NewRepository(s.DB)
	service := srv.NewService(repo)
	controller := con.NewController(service, s.validator)

	s.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Odoo Server Already Run :)")
	})

	apiv1 := s.App.Group("/api/v1")
	// travel
	travelv1 := apiv1.Group("/travel")
	travelv1.Get("", controller.GetListTravel)

}
