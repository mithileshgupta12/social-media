package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mithileshgupta12/social-media/globals"
)

type PostHandler struct{}

func (p *PostHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(globals.SuccessResponse{
		Message: "Got your posts!",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
