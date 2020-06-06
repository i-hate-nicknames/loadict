package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const dbFile = "loadict.db"

type Card struct {
	ID          uint `gorm:"primary_key"`
	Word        string
	Template    string
	Created     *time.Time
	LastFetched *time.Time
	Items       string
}

func Connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	return db
}

func LoadCards(db *gorm.DB) []*Card {
	var cards []*Card
	db.Find(&cards)
	return cards
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Card{})
}
