package server

import (
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
