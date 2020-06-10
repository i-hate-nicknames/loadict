package fetch

import (
	"log"
	"sync"

	"nvm.ga/loadict/card"
)

const concurrentFetches = 10

func FetchCards(words []string, fetcher WordFetcher) <-chan *card.Card {
	source := make(chan string, len(words))
	fetched := make(chan *Response, 0)
	rendered := make(chan *card.Card, 0)

	go fetchWords(fetcher, concurrentFetches, source, fetched)
	go renderWords(fetched, rendered)

	for _, word := range words {
		source <- word
	}
	close(source)
	return rendered
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

func renderWords(in <-chan *Response, out chan<- *card.Card) {
	for wordResponse := range in {
		log.Println("Rendering", wordResponse.Word)
		cardBack, err := renderCard(wordResponse)
		if err != nil {
			log.Println(err)
			continue
		}
		card := card.MakeCard(wordResponse.Word, cardBack)
		out <- card
	}
	close(out)
}