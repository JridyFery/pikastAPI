package router

import (
	"github.com/gorilla/mux"
	handlers "github.com/pikastAR/pikastAPI/handlers"
	//"github.com/pikastAR/pikastAPI/helpers"
)

// UserRouterHandler ...
type UserRouterHandler struct {
	Router  *mux.Router
	Handler handlers.UserHandler
}

// HandleFunctions ...
func (r *UserRouterHandler) HandleFunctions() {
	// Route Handlers / Endpoints
	r.Router.HandleFunc("/api/v1/users", r.Handler.CreateUser).Methods("POST")
}
