package controller

import (
	"Go300-backend/src/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/register", Register).Methods("POST")
	router.HandleFunc("/subscribe", Subscribe).Methods("POST")
	return router
}

func TestRegister(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Register", t, func() {
		member := &model.Member{
			Username: "MuslimBeibytuly",
		}
		jsonMember, _ := json.Marshal(member)
		request, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonMember))
		response := httptest.NewRecorder()
		Router().ServeHTTP(response, request)
		So(response.Code, ShouldEqual, 200)
	})
}
