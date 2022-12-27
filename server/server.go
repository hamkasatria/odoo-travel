package server

import (
	"fmt"
	con "odoo-travel/controllers"
	repo "odoo-travel/repositories"
	srv "odoo-travel/services"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber/v2"

	"github.com/go-playground/validator/v10"
)

type ApiServer struct {
	DB        *mongo.Database
	App       *fiber.App
	validator *validator.Validate
}

func NewServer(db *mongo.Database, validator *validator.Validate) *ApiServer {

	app := fiber.New()

	return &ApiServer{
		DB:        db,
		App:       app,
		validator: validator,
	}
}

func (s *ApiServer) ListenAndServer(port string) {
	s.registerRouters()
	s.App.Listen(":" + port)
}

func (s *ApiServer) registerRouters() {

	repo := repo.NewRepository(s.DB)
	service := srv.NewService(repo)
	controller := con.NewController(service, s.validator)

	fmt.Println(controller)

	s.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Odoo Server Already Run :)")
	})

}
