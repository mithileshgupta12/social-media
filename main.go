package main

import (
	"net/http"

	"github.com/mithileshgupta12/social-media/application"
)

func main() {
	r := application.LoadRoutes()

	http.ListenAndServe(":3000", r)
}
