package model

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubscription(t *testing.T) {
	Convey("Subscription", t, func() {
		location := Location{1, 2}
		direction := Direction{location, location}
		preference := Preference{"12:30", direction}
		preferences := make([]Preference, 2)
		preferences = append(preferences, preference)
		preferences = append(preferences, preference)
		subscription := Subscription{
			Token:       "",
			Preferences: preferences,
		}
		newSubscription, err := CreateSubscription(subscription)
		So(len(newSubscription.Token), ShouldEqual, 24)
		So(err, ShouldEqual, nil)
	})
}
