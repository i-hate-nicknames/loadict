package main

import (
	"encoding/csv"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

const fileName = "cards.csv"
const concurrentFetches = 10

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	appID, appKey := os.Getenv("APP_ID"), os.Getenv("APP_KEY")
	if appID == "" || appKey == "" {
		log.Fatal("Provide app id and app key in .env file")
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Cannot create file")
	}

	// todo: read from arguments
	words := []string{"object"}

	source := make(chan string, len(words))
	fetched := make(chan *Response, 0)
	rendered := make(chan *ExportCard, 0)

	fetcher := MakeFetcher(appID, appKey)
	go fetchWords(fetcher, concurrentFetches, source, fetched)
	go renderWords(fetched, rendered)

	for _, word := range words {
		source <- word
	}
	close(source)

	writer := csv.NewWriter(file)

	for card := range rendered {
		err := writer.Write([]string{card.word, card.card})
		if err != nil {
			log.Println(err)
		}
	}
	writer.Flush()

}

func fetchWords(fetcher WordFetcher, concurrency int, in <-chan string, out chan<- *Response) {
	var wg sync.WaitGroup
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			for word := range in {
				log.Println("Fetching ", word)
				response, err := fetcher(word)
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
		log.Println("Rendering", wordResponse.Word)
		card, err := renderCard(wordResponse)
		if err != nil {
			log.Println(err)
			continue
		}
		out <- &ExportCard{word: wordResponse.Word, card: card}
	}
	close(out)
}

type ExportCard struct {
	word, card string
}
