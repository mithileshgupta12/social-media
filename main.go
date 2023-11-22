package main

import (
	"github.com/mithileshgupta12/social-media/application"
)

func main() {
	router := &application.Router{}

	e := router.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
