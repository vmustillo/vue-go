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

func GetRoutes() {

}