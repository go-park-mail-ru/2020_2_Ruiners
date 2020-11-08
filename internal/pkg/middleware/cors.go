package middleware

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func enableCORS(cfg *models.CORSConfig, handler http.Handler) http.Handler {
	var (
		allowedOrigins = handlers.AllowedOrigins(cfg.AllowedOrigins)
		allowedHeaders = handlers.AllowedHeaders(cfg.AllowedHeaders)
		exposedHeaders = handlers.ExposedHeaders(cfg.ExposedHeaders)
		allowedMethods = handlers.AllowedMethods(cfg.AllowedMethods)
		credentials    = handlers.AllowCredentials()
	)

	return handlers.CORS(allowedOrigins, allowedHeaders, exposedHeaders, allowedMethods, credentials)(handler)
}

func CORSMiddleware(CORS models.CORSConfig) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return enableCORS(&CORS, next)
	}
}
