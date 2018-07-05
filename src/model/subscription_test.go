package model

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateSubscription(t *testing.T) {

	Convey("Subscription", t, func() {
		var preferences []Preference
		preferences = append(
			preferences,
			Preference{
				"9:30",
				Direction{
					Location{1, 2},
					Location{3, 4},
				},
			},
		)
		preferences = append(
			preferences,
			Preference{
				"18:00",
				Direction{
					Location{3, 4},
					Location{1, 2},
				},
			},
		)
		subscription := Subscription{
			Member:      Member{Username: "MuslimBeibytuly"},
			Preferences: preferences,
		}
		newSubscription, err := CreateSubscription(subscription)
		So(len(newSubscription.Id), ShouldEqual, 12)
		So(err, ShouldEqual, nil)
	})
}

func TestCreateSubscription2(t *testing.T) {

	Convey("Empty Subscription Member Username", t, func() {
		var preferences []Preference
		preferences = append(
			preferences,
			Preference{
				"9:30",
				Direction{
					Location{1, 2},
					Location{3, 4},
				},
			},
		)
		preferences = append(
			preferences,
			Preference{
				"18:00",
				Direction{
					Location{3, 4},
					Location{1, 2},
				},
			},
		)
		subscription := Subscription{
			Member:      Member{Username: ""},
			Preferences: preferences,
		}
		_, err := CreateSubscription(subscription)
		So(err, ShouldNotEqual, nil)
	})
}
