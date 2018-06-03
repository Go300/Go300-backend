package model

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"context"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"errors"
)

type Member struct {
	Token string `json:"Token"`
	Name  string `json:"Name"`
}

func Create(member Member) (Member, error) {
	if (len(member.Name) == 0) {
		return member, errors.New("Username is empty.")
	}
	client, _ := mongo.NewClient("mongodb://Go300:default@mongodb:27017")
	client.Connect(context.TODO())
	collection := client.Database("Go300DB").Collection("members")
	res, _ := collection.InsertOne(context.Background(), member)
	if oid, ok := res.InsertedID.(objectid.ObjectID); ok {
		member.Token = oid.Hex()
	}
	return member, nil
}