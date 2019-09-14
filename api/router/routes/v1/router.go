package v1

import (
	"net/http"
	"log"

	Routes "github.com/vmustillo/vue-go/api/router/routes"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hit v1 middleware")
		next.ServeHTTP(w, r)
	})
}

func GetRoutes() map[string]Routes.SubRoutePackage {
	SubRoute := map[string]Routes.SubRoutePackage{
		"/v1": {
			Routes: Routes.Routes{
				Routes.Route{
					Name: "V1HealthRoute",
					Method: "GET",
					Pattern: "/health",
					HandlerFunc: Health(),
				},
			},
			Middleware: Middleware,
		},
	}

	return SubRoute
}