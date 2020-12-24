package main

import (
	"flag"
	"github.com/Arkadiyche/http-rest-api/internal/app/apiserver"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "/home/ubuntu/back/2020_2_Ruiners/bin/apiserve", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}


//config := apiserver.Config{
//BindAddr: ":8000",
//LogLevel: "debug",
//Store:    &store.Config{DatabaseURL: "root:password@/kino_park"},
//CORS:     models.CORSConfig{},
//}