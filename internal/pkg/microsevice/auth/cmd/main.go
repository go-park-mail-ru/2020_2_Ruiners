package main

import (
	"flag"
	"github.com/Arkadiyche/http-rest-api/internal/app/apiserver"
	auth "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/auth/server"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/store"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/auth.toml", "path to config file")
}



func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	db := store.New(config.Store)
	if err != nil {
		log.Fatal(err.Error())
	}
	srv1 := auth.NewServer(config.BindAddr, db)
	if err = srv1.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
