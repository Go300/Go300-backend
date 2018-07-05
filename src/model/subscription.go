package model

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-bongo/bongo"
)

type Subscription struct {
	bongo.DocumentBase `bson:",inline"`
	Member             Member
	Preferences        []Preference
}

type Preference struct {
	Timestamp string
	Where     Direction
}

type Direction struct {
	From Location
	To   Location
}

type Location struct {
	Lat  float64
	Long float64
}

func CreateSubscription(subscription Subscription) (Subscription, error) {
	if 0 == len(subscription.Member.Username) {
		return subscription, errors.New("Member Username is empty")
	}
	config := &bongo.Config{
		ConnectionString: "Go300:default@mongodb",
		Database:         "Go300DB",
	}
	connection, err := bongo.Connect(config)

	if err != nil {
		log.Fatal(err)
	}

	err = connection.Collection("subscriptions").Save(&subscription)
	if err != nil {
		if vErr, ok := err.(*bongo.ValidationError); ok {
			fmt.Println("Validation errors are:", vErr.Errors)
		} else {
			fmt.Println("Got a real error:", err.Error())
		}
	}
	return subscription, err
}
