package model

import (
	"context"
	"errors"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Member struct {
	Token string `json:"Token"`
	Name  string `json:"Name"`
}

func Create(member Member) (Member, error) {
	if 0 == len(member.Name) {
		return member, errors.New("username is empty")
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
