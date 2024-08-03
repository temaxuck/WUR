package models

import (
	"github.com/temaxuck/WUR/service.ebooks/internal/constants"
	"gorm.io/gorm"
)

type BookFile struct {
	gorm.Model
	BookMetaID uint
	BookMeta   BookMeta `gorm:"foreignkey:BookMetaID"`
	Filename   string
	// See "github.com/temaxuck/WUR/service.ebooks/internal/db/postgres.go" Migrate() function for book_file_format type definition
	FileFormat constants.BookFileFormat `gorm:"type:book_file_format"`
}
