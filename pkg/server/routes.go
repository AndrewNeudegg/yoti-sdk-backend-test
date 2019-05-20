package server

import (
	"fmt"
	"net/http"

	"github.com/AndrewNeudegg/yoti-sdk-backend-test/pkg/logging"
	"github.com/gorilla/mux"
)

// Route is a HTTP Route definition.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of HTTP Routes.
type Routes []Route

// NewRouter compiles a new HTTP Router using the values declared in routes.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
		logging.Info(fmt.Sprintf("Added new route: %s:%s", route.Method, route.Pattern))
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"POST",
		"/",
		IndexPost,
	},
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
}
