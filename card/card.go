package card

import "time"

type Card struct {
	ID          uint `gorm:"primary_key"`
	Word        string
	Back    string
	Created     time.Time
	Updated time.Time
}

func MakeCard(word, back string) *Card {
	return &Card{Word: word, Back: back, Created: time.Now(), Updated: time.Now() }
}