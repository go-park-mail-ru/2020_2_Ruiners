package Handlers

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"../Models"
)

func enableCORS(cfg *Models.CORSConfig, handler http.Handler) http.Handler {
	var (
		allowedOrigins = handlers.AllowedOrigins(cfg.AllowedOrigins)
		allowedHeaders = handlers.AllowedHeaders(cfg.AllowedHeaders)
		exposedHeaders = handlers.ExposedHeaders(cfg.ExposedHeaders)
		allowedMethods = handlers.AllowedMethods(cfg.AllowedMethods)
		credentials = handlers.AllowCredentials()
	)

	return handlers.CORS(allowedOrigins, allowedHeaders, exposedHeaders, allowedMethods, credentials)(handler)
}
func CORSMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return enableCORS(&Models.GlobalCORSConfig, next)
	}
}