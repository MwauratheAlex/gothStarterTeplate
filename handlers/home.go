package handlers

import (
	"github.com/mwaurathealex/gothstarter/views/home"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}
