package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Unfield/cascade"
	"github.com/Unfield/zugang/api"
	"github.com/Unfield/zugang/config"
)

func main() {
	cfg := config.ZugangConfig{}

	loader := cascade.NewLoader(
		cascade.WithFile("config.toml"),
	)

	if err := loader.Load(&cfg); err != nil {
		log.Fatal(err)
	}

	apiRouter := api.NewRouter(&cfg)

	log.Printf("Serving Zugang at %s:%d", cfg.Server.Host, cfg.Server.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port), apiRouter); err != nil {
		panic(err)
	}
}
