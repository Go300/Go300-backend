package main

import (
	"Go300-backend/src/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robfig/cron"
)

func startapp() {
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.Register)
	http.Handle("/register", r)
	fmt.Println("Starting up on 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func startcron() {
	c := cron.New()
	c.AddFunc("@every 5s", func() { fmt.Println("lol") })
	c.Start()
}

func main() {
	go startcron()
	startapp()
}
