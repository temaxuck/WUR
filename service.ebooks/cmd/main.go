package main

import (
	"log"

	"github.com/temaxuck/WUR/service.ebooks/cmd/app"
	"github.com/temaxuck/WUR/service.ebooks/config"
)

func main() {
	cfg, err := config.GetConfig()

	if err != nil {
		log.Fatalln(err)
	}

	app.Run(cfg)
}
