package models

import (
	"time"
)

type BookMeta struct {
	ID              uint   `gorm:"PrimaryKey"`
	Title           string `gorm:"index"`
	Description     string
	PublicationDate time.Time
	AuthorID        uint
	Author          Author  `gorm:"foreignkey:AuthorID"`
	Tags            []Tag   `gorm:"many2many:book_meta_tags"`
	Genres          []Genre `gorm:"many2many:book_meta_genres"`
	Cover           string
}
