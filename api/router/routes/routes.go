package routes

import (
	"net/http"
	"log"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside middleware")
		next.ServeHTTP(w, r)
	})
}

func GetRoutes() Routes {
	return Routes{
		Route{
			Name: "HealthCheck",
			Method: "GET",
			Pattern: "/health",
			HandlerFunc: Health(),
		},
	}
}