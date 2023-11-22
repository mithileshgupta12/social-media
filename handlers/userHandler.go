package handlers

import (
	"net/http"
	"os"

	"github.com/mithileshgupta12/social-media/config"
)

type UserHandler struct{}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Save userId to session
	sessionConfig := config.Session{}

	cookieStore := sessionConfig.Init()
	session, err := cookieStore.Get(r, os.Getenv("SESSION_NAME"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["userId"] = 1

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
