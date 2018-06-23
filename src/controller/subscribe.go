package controller

import (
	"Go300-backend/src/model"
	"encoding/json"
	"net/http"
)

//Register member
func Subscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var subscription model.Subscription
		json.NewDecoder(r.Body).Decode(&subscription)
		newSubscription, _ := model.CreateSubscription(subscription)
		json.NewEncoder(w).Encode(newSubscription)
	}
}
