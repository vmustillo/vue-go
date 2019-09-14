package router

import (
	"net/http"

	"github.com/gorilla/mux"
	Routes "github.com/vmustillo/vue-go/api/router/routes"
)

const (
	staticDir = "/static/"
)

type RouteHandler struct {
	Router *mux.Router
}

func NewRouter() *RouteHandler {
	var router RouteHandler

	router.Router = mux.NewRouter().StrictSlash(true)

	router.Router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("." + staticDir))))

	router.Router.Use(Routes.Middleware)

	routes := Routes.GetRoutes()

	for _, route := range routes {
		router.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	
	return &router
}