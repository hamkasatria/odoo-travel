package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

type Config struct {
	DbUsername string
	DbPassword string
	DbHost     string
	DbName     string
	DbPort     string
	SvPort     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	svPort := os.Getenv("SERVER_PORT")

	return &Config{
		DbUsername: dbUsername,
		DbPassword: dbPassword,
		DbHost:     dbHost,
		DbName:     dbName,
		DbPort:     dbPort,
		SvPort:     svPort,
	}
}

func MongoDb(c *Config) (*mongo.Database, error) {
	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s", c.DbUsername, c.DbPassword, c.DbHost)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(c.DbName)

	if err != nil {
		return nil, err
	}

	mg := MongoInstance{
		Client: client,
		Db:     db,
	}

	return mg.Db, nil
}
