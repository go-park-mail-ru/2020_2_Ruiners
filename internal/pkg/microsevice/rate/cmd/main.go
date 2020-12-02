package main

import (
	"flag"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/app/apiserver"
	rate "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/rate/server"
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
	fmt.Println(db.Config())
	srv := rate.NewServer(config.BindAddr, db)
	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
