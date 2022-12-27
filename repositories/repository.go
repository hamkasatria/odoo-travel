package repositories

import (
	"odoo-travel/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetListTravels(ctx *fiber.Ctx) (*[]models.Travel, error)
	GetTravelById(ctx *fiber.Ctx, id string) (*models.Travel, error)
	AddTravel(ctx *fiber.Ctx, travel *models.Travel) error
	EditTravel(ctx *fiber.Ctx, id string, data interface{}) error
	DeleteTravel(ctx *fiber.Ctx, id string) error
}

type repository struct {
	DB *mongo.Database
}

func NewRepository(db *mongo.Database) *repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) GetListTravels(ctx *fiber.Ctx) (*[]models.Travel, error) {

	query := bson.D{{}}
	cursor, err := r.DB.Collection("travels").Find(ctx.Context(), query)
	if err != nil {
		return nil, err
	}

	var travels []models.Travel
	if err := cursor.All(ctx.Context(), &travels); err != nil {
		return nil, err
	}

	return &travels, nil
}

func (r *repository) GetTravelById(ctx *fiber.Ctx, id string) (*models.Travel, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var travel models.Travel
	query := bson.M{"_id": objectId}

	r.DB.Collection("travels").FindOne(ctx.Context(), query).Decode(&travel)

	return &travel, nil
}

func (r *repository) AddTravel(ctx *fiber.Ctx, travel *models.Travel) error {

	if _, err := r.DB.Collection("travels").InsertOne(ctx.Context(), *travel); err != nil {
		return err
	}

	return nil
}

func (r *repository) EditTravel(ctx *fiber.Ctx, id string, data interface{}) error {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	query := bson.M{"_id": objectId}

	if result := r.DB.Collection("travels").FindOneAndUpdate(ctx.Context(), query, data); result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (r *repository) DeleteTravel(ctx *fiber.Ctx, id string) error {

	objectId, err := primitive.ObjectIDFromHex(id); 
	if err != nil {
		return err
	}
	query := bson.M{"_id": objectId}

	if _, err := r.DB.Collection("travels").DeleteOne(ctx.Context(), query); err != nil {
		return err
	}
	return nil
}
