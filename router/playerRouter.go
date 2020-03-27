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
	r.Router.HandleFunc("/api/v1/register", r.Handler.CreatePlayer).Methods("POST")
	r.Router.HandleFunc("/api/v1/login", r.Handler.Login).Methods("GET")
	r.Router.Handle("/api/v1/player", helpers.IsAuthorized(r.Handler.GetPlayer)).Methods("GET")
	r.Router.Handle("/api/v1/players", helpers.IsAuthorized(r.Handler.GetPlayers)).Methods("GET")
	r.Router.Handle("/api/v1/playerpokemon", helpers.IsAuthorized(r.Handler.AddPokemonPlayer)).Methods("POST")
	r.Router.Handle("/api/v1/playerpokemons", helpers.IsAuthorized(r.Handler.GetPlayerPokemons)).Methods("GET")
	r.Router.Handle("/api/v1/playerby", helpers.IsAuthorized(r.Handler.GetPlayerBy)).Methods("GET")
	r.Router.Handle("/api/v1/player", helpers.IsAuthorized(r.Handler.DeletePlayer)).Methods("DELETE")
	r.Router.Handle("/api/v1/player", helpers.IsAuthorized(r.Handler.UpdatePlayer)).Methods("PUT")
	//	r.Router.Handle("/api/v1/updateplayerpic", helpers.IsAuthorized(r.Handler.UpdatePlayerPic)).Methods("POST")
	r.Router.HandleFunc("/api/v1/updateplayerpic", r.Handler.UpdatePlayerPic).Methods("POST")
	//64bits => byte => file for pics
}
