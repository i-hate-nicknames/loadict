package card

import (
	"strings"
	"time"
)

type Card struct {
	ID       uint `gorm:"primary_key"`
	Word     string
	Back     string
	Created  time.Time
	Updated  time.Time
	Exported bool
}

func MakeCard(word, back string) *Card {
	return &Card{
		Word:    strings.ToLower(word),
		Back:    back,
		Created: time.Now(),
		Updated: time.Now(),
	}
}
