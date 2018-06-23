package main

import (
	"Go300-backend/src/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func startapp() {
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.Register)
	r.HandleFunc("/subscribe", controller.Subscribe)
	http.Handle("/", r)
	fmt.Println("Starting up on 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	startapp()
}
