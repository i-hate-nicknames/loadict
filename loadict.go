package main

import (
	"encoding/csv"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

const fileName = "cards.csv"

func main() {

	source := make(chan string, 0)
	fetched := make(chan *Response, 0)
	rendered := make(chan *ExportCard, 0)

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Cannot create file")
	}
	go fetchWords(10, source, fetched)
	go renderWords(fetched, rendered)

	words := []string{"entail", "whirlwind", "smart", "entail", "whirlwind", "smart"}
	// we have to put words into source in a different goroutine because initial capacity
	// is 0, and we may get blocked by put
	// Alternatively we could've set channel capacity to the number of words
	go func() {
		for _, word := range words {
			source <- word
		}
		close(source)
	}()

	writer := csv.NewWriter(file)

	for card := range rendered {
		err := writer.Write([]string{card.word, card.card})
		if err != nil {
			log.Println(err)
		}
	}
	writer.Flush()

}

func fetchWords(concurrency int, in <-chan string, out chan<- *Response) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	appID, appKey := os.Getenv("APP_ID"), os.Getenv("APP_KEY")
	if appID == "" || appKey == "" {
		log.Fatal("Provide app id and app key in .env file")
	}
	var wg sync.WaitGroup
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			for word := range in {
				log.Println("Fetching ", word)
				response, err := fetchWord(appID, appKey, word)
				if err != nil {
					log.Println(err)
					continue
				}
				out <- response
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(out)
}

func renderWords(in <-chan *Response, out chan<- *ExportCard) {
	for wordResponse := range in {
		log.Println("Rendering", wordResponse.word)
		card, err := renderCard(wordResponse)
		if err != nil {
			log.Println(err)
			continue
		}
		out <- &ExportCard{word: wordResponse.word, card: card}
	}
	close(out)
}

type ExportCard struct {
	word, card string
}
