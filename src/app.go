package main

import (
	"Go300-backend/src/controller"
	"fmt"
	"log"
	"net/http"

	"Go300-backend/src/jobs"

	"strconv"

	"github.com/gorilla/mux"
	"github.com/jasonlvhit/gocron"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func startapp() {
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.Register)
	r.HandleFunc("/subscribe", controller.Subscribe)
	http.Handle("/", r)
	fmt.Println("Starting up on 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func startcron() {
	fmt.Println("Setting up cron jobs")
	for hour := 0; hour < 24; hour++ {
		h := strconv.Itoa(hour)
		gocron.Every(1).Day().At(h + ":00").Do(jobs.ConfirmationService)
		gocron.Every(1).Day().At(h + ":30").Do(jobs.ConfirmationService)
	}
	<-gocron.Start()
}

func main() {
	go startcron()
	startapp()
}
