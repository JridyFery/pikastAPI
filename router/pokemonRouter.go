package router

import (
	"github.com/gorilla/mux"
	handlers "github.com/JridyFery/pikastAPI/handlers"
	"github.com/JridyFery/pikastAPI/helpers"
)

// PokemonRouterHandler ...
type PokemonRouterHandler struct {
	Router  *mux.Router
	Handler handlers.PokemonHandler
}

// HandleFunctions ...
func (r *PokemonRouterHandler) HandleFunctions() {
	// Route Handlers / Endpoints
	r.Router.Handle("/api/v1/pokemon", helpers.IsAuthorized(r.Handler.CreatePokemon)).Methods("POST")
	//r.Router.HandleFunc("/api/v1/pokemon", (r.Handler.CreatePokemon)).Methods("POST")
	r.Router.Handle("/api/v1/pokemon", helpers.IsAuthorized(r.Handler.GetPokemon)).Methods("GET")
	r.Router.Handle("/api/v1/pokemons", helpers.IsAuthorized(r.Handler.GetPokemons)).Methods("GET")
	r.Router.Handle("/api/v1/pokemon", helpers.IsAuthorized(r.Handler.DeletePokemon)).Methods("DELETE")
	r.Router.Handle("/api/v1/pokemon", helpers.IsAuthorized(r.Handler.UpdatePokemon)).Methods("PUT")
}
