package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route structs for routes
type Route struct {
	URI          string
	Method       string
	FunctionName func(http.ResponseWriter, *http.Request)
}

// ConfigureRoute configures application routes
func ConfigureRoute(router *mux.Router) *mux.Router {

	routes := bookRouters

	for _, route := range routes {
		router.HandleFunc(route.URI, route.FunctionName).Methods(route.Method)
	}

	return router
}
