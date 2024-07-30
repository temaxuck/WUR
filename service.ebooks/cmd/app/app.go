package app

import (
	"github.com/gin-gonic/gin"
	"github.com/temaxuck/WUR/service.ebooks/config"
	"github.com/temaxuck/WUR/service.ebooks/internal/db"
	"github.com/temaxuck/WUR/service.ebooks/internal/rest"
)

func Run(cfg config.Config) {
	router := gin.Default()
	dbHandler := db.Init(cfg.PostgresURL)

	rest.RegisterRoutes(router, dbHandler)

	router.Run(cfg.GetAddress())
}
