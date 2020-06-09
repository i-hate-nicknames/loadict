package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"nvm.ga/loadict/card"
)

const dbFile = "loadict.db"

func Connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	return db
}

func LoadCards(db *gorm.DB) []*card.Card {
	var cards []*card.Card
	db.Find(&cards)
	return cards
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&card.Card{})
}
