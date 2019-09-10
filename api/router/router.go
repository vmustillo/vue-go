package router

import (
	"net/http"

	"github.com/gorilla/mux"
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

	return &router
}