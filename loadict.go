package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const fileName = "cards.csv"

func main() {

	source := make(chan string, 0)
	fetched := make(chan *Response, 0)
	errors := make(chan error, 0)
	rendered := make(chan *ExportCard, 0)
	results := make(chan bool, 0)

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Cannot create file")
	}
	go fetchWords(1, source, fetched, errors)
	go renderWords(fetched, rendered, errors)
	go exportWords(file, rendered, results, errors)

	words := []string{"entail", "whirlwind", "smart"}
	for _, word := range words {
		source <- word
	}
	close(source)

	n := 0
	for n < len(words) {
		select {
		case <-results:
			log.Println("exported a word!")
		case err := <-errors:
			log.Println(err)
		}
		n++
	}

}

func fetchWords(parallelism int, in <-chan string, out chan<- *Response, errors chan<- error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	appID, appKey := os.Getenv("APP_ID"), os.Getenv("APP_KEY")
	if appID == "" || appKey == "" {
		log.Fatal("Provide app id and app key in .env file")
	}
	for word := range in {
		log.Println("Fetching ", word)
		response, err := fetchWord(appID, appKey, word)
		if err != nil {
			errors <- err
			continue
		}
		out <- response
	}
	// var wg sync.WaitGroup
	// wg.Add(parallelism)
	// for i := 0; i < parallelism; i++ {
	// 	go func() {
	// 		for word := range in {
	// 			log.Println("Fetching ", word)
	// 			response, err := fetchWord(appID, appKey, word)
	// 			if err != nil {
	// 				errors <- err
	// 				continue
	// 			}
	// 			out <- response
	// 		}
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()
	close(out)
}

func renderWords(in <-chan *Response, out chan<- *ExportCard, errors chan<- error) {
	for wordResponse := range in {
		log.Println("Rendering", wordResponse.Results[0].Word)
		card, err := renderCard(wordResponse)
		if err != nil {
			errors <- err
			continue
		}
		out <- &ExportCard{word: wordResponse.Results[0].Word, card: card}
	}
	close(out)
}

func exportWords(writer io.Writer, in <-chan *ExportCard, results chan<- bool, errors chan<- error) {
	for card := range in {
		log.Println("Exporting ", card.word)
		err := export(writer, card)
		if err != nil {
			errors <- err
			continue
		}
		results <- true
	}
	close(results)
	close(errors)
}

type ExportCard struct {
	word, card string
}

func export(w io.Writer, card *ExportCard) error {
	writer := csv.NewWriter(w)
	err := writer.Write([]string{card.word, card.card})
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}
