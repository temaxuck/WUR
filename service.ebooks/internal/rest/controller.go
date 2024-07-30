package rest

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/temaxuck/WUR/service.ebooks/internal/rest/handlers"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := handlers.Handler{
		DB: db,
	}

	booksRoutes := router.Group("/books")
	booksRoutes.POST("/", h.CreateBook)
	booksRoutes.GET("/", h.GetBooks)
	booksRoutes.GET("/:id", h.GetBook)
	booksRoutes.PUT("/:id", h.UpdateBook)
	booksRoutes.DELETE("/:id", h.DeleteBook)

	authorsRoutes := router.Group("/authors")
	authorsRoutes.POST("/", h.CreateAuthor)
	authorsRoutes.GET("/", h.GetAuthors)
	authorsRoutes.GET("/:id", h.GetAuthor)
	authorsRoutes.PUT("/:id", h.UpdateAuthor)
	authorsRoutes.DELETE("/:id", h.DeleteAuthor)
}
