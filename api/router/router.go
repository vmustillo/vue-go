package router

import (
	"net/http"

	"github.com/gorilla/mux"

	Routes "github.com/vmustillo/vue-go/api/router/routes"
	V1Routes "github.com/vmustillo/vue-go/api/router/routes/v1"
)

const (
	staticDir = "/static/"
)

type RouteHandler struct {
	Router *mux.Router
}

func (r *RouteHandler) AttachSubRouterWithMiddleware(path string, subroutes Routes.Routes, Middleware mux.MiddlewareFunc) (*mux.Router){
	SubRouter := r.Router.PathPrefix(path).Subrouter()
	SubRouter.Use(Middleware)

	for _, route := range subroutes {
		SubRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return SubRouter
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

	v1SubRoutes := V1Routes.GetRoutes()
	
	for name, pack := range v1SubRoutes {
		router.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}
	
	return &router
}