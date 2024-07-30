package models

import "time"

type Author struct {
	ID           uint   `gorm:"PrimaryKey"`
	FullName     string `gorm:"index"`
	BirthDate    time.Time
	DeathDate    time.Time
	Description  string
	WikipediaURL string
	Image        string
}

func GetDefaultAuthor() Author {
	return Author{
		FullName:    "Unknown",
		Description: "Unknown author",
	}
}
