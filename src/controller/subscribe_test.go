package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"Go300-backend/src/model"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubscribe(t *testing.T) {
	Convey("Subscribe", t, func() {
		location := model.Location{1, 2}
		direction := model.Direction{location, location}
		preference := model.Preference{"12:30", direction}
		preferences := make([]model.Preference, 2)
		preferences = append(preferences, preference)
		preferences = append(preferences, preference)
		subscription := model.Subscription{
			Preferences: preferences,
		}
		jsonSubscription, _ := json.Marshal(subscription)
		request, _ := http.NewRequest("POST", "/subscribe", bytes.NewBuffer(jsonSubscription))
		response := httptest.NewRecorder()
		Router().ServeHTTP(response, request)
		So(response.Code, ShouldEqual, 200)
	})
}
