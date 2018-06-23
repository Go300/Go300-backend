package controller

import (
	"Go300-backend/src/model"
	"encoding/json"
	"net/http"
)

//Register member
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var member model.Member
		json.NewDecoder(r.Body).Decode(&member)
		newMember, _ := model.CreateMember(member)
		json.NewEncoder(w).Encode(newMember)
	}
}
