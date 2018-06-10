package main

import (
	"Go300-backend/src/controller"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.Register)
	http.Handle("/register", r)
	fmt.Println("Starting up on 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
