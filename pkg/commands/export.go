package commands

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"nvm.ga/loadict/pkg/db"
)

const exportFile = "cards.csv"

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.PersistentFlags().IntVarP(&exportNum, "number", "n", 10, "number of entries to export")
}

var exportNum int

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export cards to an anki deck",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := db.GetDB()
		if err != nil {
			log.Fatalf("cannot connect to db: %s", err)
		}
		err = exportCards(exportNum, db)
		if err != nil {
			log.Fatalf("cannot export cards: %s", err)
		}
	},
}

// Export at most num cards from the db into .csv file, ready
// to be imported by anki
// Exported cards will be marked as such in the db and won't be
// exported in future
func exportCards(num int, conn *gorm.DB) error {
	file, err := os.Create(exportFile)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)

	cards, err := db.LoadCards(conn, num)
	if err != nil {
		return err
	}

	if len(cards) == 0 {
		log.Println("There are no cards to be exported: load more words")
		return nil
	}

	fmt.Printf("Exporting %d words to %s\n", len(cards), exportFile)
	for _, card := range cards {
		err := writer.Write([]string{card.Word, card.Back})
		if err != nil {
			return err
		}
		card.Exported = true
	}
	writer.Flush()
	return db.SaveCards(conn, cards)
}
