package router

import (
	"api/src/router/route"

	"github.com/gorilla/mux"
)

// Create Resposible for create application routes
func Create() *mux.Router {
	router := mux.NewRouter()
	return route.ConfigureRoute(router)

}
