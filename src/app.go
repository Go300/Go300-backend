package main

import (
	"Go300-backend/src/controller"
	"fmt"
	"log"
	"net/http"

	"os"

	"Go300-backend/src/jobs"

	"github.com/gorilla/mux"
	"github.com/robfig/cron"
)

func startapp() {
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.Register)
	r.HandleFunc("/subscribe", controller.Subscribe)
	http.Handle("/", r)
	fmt.Println("Starting up on 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func startcron() {
	c := cron.New()
	c.AddFunc(fmt.Sprintf("@every %s", os.Getenv("CONFIRMATION_CRON_TIMER")), func() { jobs.ConfirmationService() })
	c.Start()
}

func main() {
	go startcron()
	startapp()
}
