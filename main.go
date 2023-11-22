package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mithileshgupta12/social-media/application"
	"github.com/mithileshgupta12/social-media/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	database := config.Database{}
	if err := database.Init(); err != nil {
		log.Fatal(err.Error())
	}

	router := &application.Router{}

	r := router.Init()

	http.ListenAndServe(":8000", r)
}
