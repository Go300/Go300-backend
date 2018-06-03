package controller

import (
	"net/http"
	"encoding/json"
	"Go300/src/model"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "POST") {
		var member model.Member
		json.NewDecoder(r.Body).Decode(&member)
		new_member, _ := model.Create(member)
		json.NewEncoder(w).Encode(new_member)
	}
}