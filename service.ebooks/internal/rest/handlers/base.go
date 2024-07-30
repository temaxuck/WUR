package handlers

import (
	"github.com/temaxuck/WUR/service.ebooks/pkg/models"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func fetchDefaultAuthor(db *gorm.DB) (models.Author, error) {
	var author models.Author
	result := db.Model(&models.Author{FullName: "Unknown"}).First(&author)
	return author, result.Error
}
