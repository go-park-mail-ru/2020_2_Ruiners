package main

import (
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/app/apiserver"
	auth "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/auth/server"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/store"
	"github.com/BurntSushi/toml"
	"log"
)


func main() {

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile("/home/ubuntu/back/2020_2_Ruiners/configs/auth.toml", config)
	if err != nil {
		config = &apiserver.Config{
			BindAddr: ":8001",
			LogLevel: "debug",
			Store:    &store.Config{DatabaseURL: "root:password@/kino_park"},
			CORS:     models.CORSConfig{},
		}
	}
	db := store.New(config.Store)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(db.Config())
	srv := auth.NewServer(config.BindAddr, db)
	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
