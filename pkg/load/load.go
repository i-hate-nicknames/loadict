package load

import (
	"log"
	"sync"
	"time"

	"nvm.ga/loadict/pkg/card"
	"nvm.ga/loadict/pkg/load/loader"
)

const concurrentFetches = 10

type loadResult struct {
	Word, Result string
}

func FetchCards(words []string, l loader.Loader) []*card.Card {
	wordsChan := make(chan string, len(words))
	results := make(chan loadResult, 0)

	go fetchWords(l, concurrentFetches, wordsChan, results)

	cardsPut := 0
	for _, word := range words {
		if cardsPut%l.GetRPM() == 0 && cardsPut > 0 {
			time.Sleep(1 * time.Minute)
		}
		wordsChan <- word
		cardsPut++
	}
	close(wordsChan)

	cards := make([]*card.Card, 0)
	for result := range results {
		log.Println("Fetched", result.Word)
		cards = append(cards, card.MakeCard(result.Word, result.Result))
	}
	return cards
}

func fetchWords(l loader.Loader, concurrency int, in <-chan string, out chan<- loadResult) {
	var wg sync.WaitGroup
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			for word := range in {
				log.Println("Fetching", word)
				response, err := l.Load(word)
				if err != nil {
					log.Println(err)
					continue
				}
				out <- loadResult{word, response}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(out)
}
