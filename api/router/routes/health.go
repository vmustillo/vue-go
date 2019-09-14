package routes

import (
	"net/http"
)

func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is healthy!"))
	}
}