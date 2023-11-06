package database

import (
	"fmt"
	"log"
	"os"
	"time"
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go/mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client{
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error locading the .env file")
	}

	MongoDb := os.Getenv("THE_MONGODB_URL")
}