package models

import "gorm.io/gorm"

type BookFileFormat string

/*
When adding a new enum value here, don't forget to add it to the
"github.com/temaxuck/WUR/service.ebooks/internal/db/postgres.go"'s
Migrate function.

TODO: Find a way to fix this inconsistency
*/
const (
	LRF    BookFileFormat = "LRF"
	LRX    BookFileFormat = "LRX"
	DJVU   BookFileFormat = "DJVU"
	EPUB   BookFileFormat = "EPUB"
	FB2    BookFileFormat = "FB2"
	PDF    BookFileFormat = "PDF"
	IBOOKS BookFileFormat = "IBOOKS"
	AZW    BookFileFormat = "AZW"
	MOBI   BookFileFormat = "MOBI"
	PDB    BookFileFormat = "PDB"
	TXT    BookFileFormat = "TXT"
	RTF    BookFileFormat = "RTF"
)

type BookFile struct {
	gorm.Model
	BookMetaID uint
	BookMeta   BookMeta `gorm:"foreignkey:BookMetaID"`
	Filename   string
	// See "github.com/temaxuck/WUR/service.ebooks/internal/db/postgres.go" Migrate() function for book_file_format type definition
	FileFormat BookFileFormat `gorm:"type:book_file_format"`
}
