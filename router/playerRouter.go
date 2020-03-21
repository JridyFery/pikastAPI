package router

import (
	"github.com/gorilla/mux"
	handlers "github.com/pikastAR/pikastAPI/handlers"
	"github.com/pikastAR/pikastAPI/helpers"
)

// PlayerRouterHandler ...
type PlayerRouterHandler struct {
	Router  *mux.Router
	Handler handlers.PlayerHandler
}

// HandleFunctions ...
func (r *PlayerRouterHandler) HandleFunctions() {
	// Route Handlers / Endpoints
	r.Router.HandleFunc("/api/v1/players", r.Handler.CreatePlayer).Methods("POST")
	r.Router.HandleFunc("/api/v1/player", r.Handler.GetPlayer).Methods("GET")
	r.Router.HandleFunc("/api/v1/player", r.Handler.DeletePlayer).Methods("DELETE")
	r.Router.Handle("/api/v1/player", helpers.IsAuthorized(r.Handler.UpdatePlayer)).Methods("PUT")
	r.Router.Handle("/api/v1/players", helpers.IsAuthorized(r.Handler.GetPlayers)).Methods("GET")
}
