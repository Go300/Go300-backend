package model

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Subscription struct {
	Token       string       `json:"Token"`
	Preferences []Preference `json:"Preferences"`
}

type Preference struct {
	Timestamp string    `json:"Timestamp"`
	Where     Direction `json:"Where"`
}

type Direction struct {
	From Location `json:"From"`
	To   Location `json:"To"`
}

type Location struct {
	Lat  float64 `json:"Lat"`
	Long float64 `json:"Long"`
}

func CreateSubscription(subscription Subscription) (Subscription, error) {
	client, _ := mongo.NewClient("mongodb://Go300:default@mongodb:27017")
	client.Connect(context.TODO())
	collection := client.Database("Go300DB").Collection("subscriptions")
	res, _ := collection.InsertOne(context.Background(), subscription)
	if oid, ok := res.InsertedID.(objectid.ObjectID); ok {
		subscription.Token = oid.Hex()
	}
	return subscription, nil
}
