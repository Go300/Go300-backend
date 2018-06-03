package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
	"Go300-backend/src/controller"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.Register)
	http.Handle("/register", r)
	fmt.Println("Starting up on 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
