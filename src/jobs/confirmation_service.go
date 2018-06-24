package jobs

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func ConfirmationService() {
	client, _ := mongo.NewClient("mongodb://Go300:default@mongodb:27017")
	client.Connect(context.TODO())
	collection := client.Database("Go300DB").Collection("waiting")
	collection.Name()
}
