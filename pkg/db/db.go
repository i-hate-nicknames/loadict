package db

import (
	"github.com/jinzhu/gorm"
	// sqlite extension for gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"nvm.ga/loadict/pkg/card"
)

const dbFile = "loadict.db"

func GetDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&card.Card{})
	return db, nil
}

// LoadCards loads at most num cards that have not yet been
// exported
func LoadCards(db *gorm.DB, num int) ([]*card.Card, error) {
	var cards []*card.Card
	if res := db.Where("exported = ?", false).Find(&cards).Limit(num); res.Error != nil {
		return nil, res.Error
	}
	return cards, nil
}

func LoadCardsByWords(db *gorm.DB, words []string) ([]*card.Card, error) {
	var cards []*card.Card
	if res := db.Where("word IN (?)", words).Find(&cards); res.Error != nil {
		return nil, res.Error
	}
	return cards, nil
}

func SaveCards(db *gorm.DB, cards []*card.Card) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		for _, card := range cards {
			tx.Save(card)
		}
		return nil
	})
	return err
}
