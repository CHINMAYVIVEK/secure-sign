package server

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Import to enable profiling
	"secure-sign/app/health"
	"secure-sign/app/user"
	"secure-sign/config"

	"github.com/gorilla/mux"
)

// APIVersion represents the version of the API.
const APIVersion = "v1"

// NewRouter creates a new router with all routes.
func NewRouter() *mux.Router {
	serverAddr := fmt.Sprintf(":%s", config.Cfg.Server.Port)

	router := mux.NewRouter()

	// Register profiling endpoints
	router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)

	log.Printf("Registered pprof endpoints. Server is running on http://localhost%s/debug/pprof/\n", serverAddr)

	// User routes with versioning
	router.HandleFunc("/"+APIVersion+"/health", health.HealthCheck).Methods(http.MethodGet)
	router.HandleFunc("/"+APIVersion+"/user/register", user.RegisterHandler).Methods(http.MethodPost)
	router.HandleFunc("/"+APIVersion+"/user/login", user.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/"+APIVersion+"/user/{username}", user.GetUserHandler).Methods(http.MethodGet)

	return router
}

// ProfileHandler is a middleware to enable profiling for specific requests.
func ProfileHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Enable profiling for this request
		profileKey := "pprof"
		r.URL.Path = "/" + profileKey + r.URL.Path
		http.DefaultServeMux.ServeHTTP(w, r)
	}
}
