package main

import (
	"github.com/Arkadiyche/http-rest-api/internal/app/apiserver"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/store"
	"log"
)



func main() {

	config := apiserver.Config{
		BindAddr: ":8000",
		LogLevel: "debug",
		Store:    &store.Config{DatabaseURL: "root:password@/kino_park"},
		CORS:     models.CORSConfig{},
	}
	s := apiserver.New(&config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
