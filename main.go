package main

import (
	"odoo-travel/config"
	"odoo-travel/server"

	"github.com/go-playground/validator/v10"
)

func main() {
	validation := validator.New()
	cfg := config.LoadConfig()
	dbInit, err := config.MongoDb(cfg)
	if err != nil {
		panic(err)
	}

	server := server.NewServer(dbInit, validation)
	server.ListenAndServer(cfg.SvPort)
}
