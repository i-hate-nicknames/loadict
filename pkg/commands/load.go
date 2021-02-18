package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"nvm.ga/loadict/pkg/card"
	"nvm.ga/loadict/pkg/db"
	"nvm.ga/loadict/pkg/load"
	"nvm.ga/loadict/pkg/load/oxford"
)

func init() {
	rootCmd.AddCommand(loadCmd)
}

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "load and process words for later export",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := db.GetDB()
		if err != nil {
			log.Fatalf("cannot connect to db: %s", err)
		}
		err = loadWords(db)
		if err != nil {
			log.Fatalf("cannot load words: %s", err)
		}
	},
}

// loadWords takes list of words from stdin, each word on its own line,
// loads definitions of these words using dictionary API, generates
// html card body using response data and saves these cards to the db
func loadWords(conn *gorm.DB) error {

	fmt.Println("loading words")
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading env file: %w", err)
	}
	appID, appKey := os.Getenv("APP_ID"), os.Getenv("APP_KEY")
	if appID == "" || appKey == "" {
		return fmt.Errorf("app id or app key is not specified in env file")
	}

	words, err := readWords()
	if err != nil {
		return err
	}
	existingCards, err := db.LoadCardsByWords(conn, words)
	if err != nil {
		return err
	}

	wordsToFetch := filterExisting(words, existingCards)

	if len(wordsToFetch) == 0 {
		log.Println("All the words are already loaded")
		return nil
	}

	log.Println("Loading the following words:", wordsToFetch)
	cards := load.FetchCards(wordsToFetch, oxford.Loader{AppID: appID, AppKey: appKey})

	fmt.Printf("Fetched %d cards!\n", len(cards))

	return db.SaveCards(conn, cards)
}

func readWords() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	words := make([]string, 0)
	for scanner.Scan() {
		words = append(words, strings.TrimSpace(scanner.Text()))
	}
	return words, scanner.Err()
}

func filterExisting(words []string, existing []*card.Card) []string {
	result := make([]string, 0)
	existingMap := make(map[string]bool, 0)
	for _, card := range existing {
		existingMap[card.Word] = true
	}
	for _, word := range words {
		if !existingMap[word] {
			result = append(result, word)
		}
	}
	return result
}
