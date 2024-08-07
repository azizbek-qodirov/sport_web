package main

import (
	"log"
	"project/api"
	"project/config"
	"project/storage"
)

func main() {
	cf := config.Load()
	dbs, err := storage.NewStorage(cf)
	if err != nil {
		log.Fatal(err)
	}
	defer dbs.Close()

	r := api.NewRouter(dbs)
	if err := r.Run(cf.GATEWAY_PORT); err != nil {
		log.Fatal(err)
	}
}
