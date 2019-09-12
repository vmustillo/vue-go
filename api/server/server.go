package server

import (
	"net/http"
	"time"
	"log"
	"strconv"
	"os"

	"github.com/gorilla/handlers"
	RouterFactory "github.com/vmustillo/vue-go/api/router"
)

// Server - struct with port and address
type Server struct {
	Port int
	Addr string
	HTTPServer *http.Server
}

// Start - starts server service
func (s *Server) Start() {
	log.Println("Server started on port", s.Port)
	log.Fatal(s.HTTPServer.ListenAndServe())
}

// NewServer - Initializes a new server
func NewServer(port int) *Server {
	var server Server

	server.Port = port
	server.Addr = ":" + strconv.Itoa(port)

	router := RouterFactory.NewRouter()

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}),
		handlers.AllowedHeaders([]string{"COntent-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(router.Router))
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	server.HTTPServer = &http.Server{
		Addr: server.Addr,
		Handler: router.Router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &server
}