package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mithileshgupta12/social-media/application"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	router := &application.Router{}

	r := router.Init()

	http.ListenAndServe(":8000", r)
}
