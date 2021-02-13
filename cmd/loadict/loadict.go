package main

import (
	"nvm.ga/loadict/pkg/commands"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	commands.Execute()
}
