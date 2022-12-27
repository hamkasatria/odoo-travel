package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
}

type repository struct {
	DB *mongo.Database
}

func NewRepository(db *mongo.Database) *repository {
	return &repository{
		DB: db,
	}
}
