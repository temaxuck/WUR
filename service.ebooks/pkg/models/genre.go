package models

type Genre struct {
	ID    uint       `gorm:"PrimaryKey"`
	Name  string     `gorm:"uniqueIndex"`
	Books []BookMeta `gorm:"many2many:book_meta_genres"`
}
