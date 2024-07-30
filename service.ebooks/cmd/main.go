package main

import (
	"log"

	"github.com/temaxuck/WUR/service.ebooks/cmd/app"
	"github.com/temaxuck/WUR/service.ebooks/config"
)

func main() {
	var cfg config.Config

	if err := cfg.LoadConfig(); err != nil {
		log.Fatalln(err)
	}

	app.Run(cfg)
}
