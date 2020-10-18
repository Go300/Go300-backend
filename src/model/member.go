package model

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-bongo/bongo"
)

//Member model
type Member struct {
	bongo.DocumentBase `bson:",inline"`
	Username           string
}

//CreateMember member
func CreateMember(member Member) (Member, error) {
	if 0 == len(member.Username) {
		return member, errors.New("Member Username is empty")
	}
	config := &bongo.Config{
		ConnectionString: fmt.Sprintf("%s:%s@%s", os.Getenv("DB_ROOT_USERNAME"), os.Getenv("DB_ROOT_PASSWORD"), os.Getenv("DB_URL")),
		Database:         os.Getenv("DB_NAME"),
	}
	connection, err := bongo.Connect(config)

	if err != nil {
		log.Fatal(err)
	}
	err = connection.Collection("members").Save(&member)
	if err != nil {
		if vErr, ok := err.(*bongo.ValidationError); ok {
			fmt.Println("Validation errors are:", vErr.Errors)
		} else {
			fmt.Println("Got a real error:", err.Error())
		}
	}
	return member, nil
}
