package config

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

type Session struct{}

func (s *Session) Init() *sessions.CookieStore {
	sessionKey := os.Getenv("SESSION_KEY")
	store := sessions.NewCookieStore([]byte(sessionKey))
	store.Options.HttpOnly = true
	store.Options.SameSite = http.SameSiteLaxMode
	store.Options.Secure = false

	return store
}
