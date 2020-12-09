package apiserver

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/store"
)

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
	CORS     models.CORSConfig
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8000",
		LogLevel: "debug",
		Store:    store.NewConfig(),
		CORS: models.CORSConfig{
			AllowedOrigins: []string{"http://localhost", "http://95.163.208.72:3000", "http://localhost:3000", "http://kino-park.online", "http://127.0.0.1"},
			AllowedHeaders: []string{"If-Modified-Since", "Cache-Control", "Content-Type", "Range"},
			AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"},
			ExposedHeaders: []string{"Content-Length", "Content-Range"},
		},
	}
}
