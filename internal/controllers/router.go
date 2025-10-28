package controllers

import (
	"dev-clash/internal/controllers/handlers"

	"github.com/gorilla/mux"
)

const (
	url_prefix = "/api/v1"
)

func InitRoter(h *handlers.Handlers) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(url_prefix+"/user/register", h.CreateUser).Methods("POST")

	return router
}
