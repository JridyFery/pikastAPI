package router

import (
	"github.com/gorilla/mux"
	handlers "github.com/pikastAR/pikastAPI/handlers"
)

// PokemonRouterHandler ...
type PokemonRouterHandler struct {
	Router  *mux.Router
	Handler handlers.PokemonHandler
}

// HandleFunctions ...
func (r *PokemonRouterHandler) HandleFunctions() {
	// Route Handlers / Endpoints
	r.Router.HandleFunc("/api/v1/pokemon", r.Handler.CreatePokemon).Methods("POST")
	r.Router.HandleFunc("/api/v1/pokemon", r.Handler.GetPokemon).Methods("GET")

}
