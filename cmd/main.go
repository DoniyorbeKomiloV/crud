package main

import (
	"fmt"
	"net/http"
	"user/api"
	"user/config"
	"user/storage/postgres"
)

func main() {

	cfg := config.Load()
	strg, err := postgres.NewConnectionPostgres(cfg)
	if err != nil {
		panic(err)
	}
	defer strg.Close()

	api.NewApi(&cfg, strg)

	fmt.Println("Listening " + cfg.HTTPPort + " port... ")
	err = http.ListenAndServe(cfg.ServerHost+cfg.HTTPPort, nil)
	if err != nil {
		panic(err)
	}
}
