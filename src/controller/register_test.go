package controller

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"Go300-backend/src/model"
	"bytes"
	"encoding/json"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/register", Register).Methods("POST")
	return router
}


func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		member := &model.Member {
			Name: "test name 1",
			Token: "",
		}
		jsonMember, _ := json.Marshal(member)
		request, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonMember))
		response := httptest.NewRecorder()
		Router().ServeHTTP(response, request)
		So(response.Code, ShouldEqual, 200)
	})
}