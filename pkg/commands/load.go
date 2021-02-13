package commands

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"nvm.ga/loadict/pkg/card"
	"nvm.ga/loadict/pkg/db"
	"nvm.ga/loadict/pkg/fetch"
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
	fetcher := fetch.MakeFetcher(appID, appKey)
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
	cards := fetch.FetchCards(wordsToFetch, fetcher)

	fmt.Printf("Fetched %d cards!\n", len(cards))

	return db.SaveCards(conn, cards)
}

func readWords() ([]string, error) {
	reader := bufio.NewReader(os.Stdin)
	words := make([]string, 0)
	for {
		word, err := reader.ReadString('\n')
		if err == io.EOF {
			return words, nil
		}
		if err != nil {
			return nil, fmt.Errorf("Cannot read words: %w", err)
		}
		words = append(words, strings.TrimSpace(word))
	}
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
