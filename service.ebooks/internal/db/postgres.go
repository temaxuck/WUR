package db

import (
	"errors"
	"log"

	"github.com/temaxuck/WUR/service.ebooks/internal/constants"
	"github.com/temaxuck/WUR/service.ebooks/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.Exec(
		`DO $$
		BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'book_file_format') THEN 
			CREATE TYPE book_file_format AS ENUM(
		` + constants.GetBookFileFormats(true) +
			`
		END IF;
		END$$;`,
	)

	return db.AutoMigrate(
		&models.BookFile{},
		&models.BookMeta{},
		&models.Author{},
		&models.Genre{},
		&models.Tag{},
	)
}

func EnsureDefaultInstances(db *gorm.DB) error {
	log.Println("Checking if default instances present in database")

	// Ensure default Author instance exists
	result := db.First(&models.Author{FullName: "Unknown"})

	if result.Error == nil {
		log.Println("Default author exists")
	} else {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("Creating default author instance")
			author := models.GetDefaultAuthor()
			tx := db.Create(&author)
			if tx.Error != nil {
				return tx.Error
			}
		} else {
			return result.Error
		}
	}

	return nil
}

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := Migrate(db); err != nil {
		log.Fatalln(err)
	}

	if err := EnsureDefaultInstances(db); err != nil {
		log.Fatalln(err)
	}

	return db
}
