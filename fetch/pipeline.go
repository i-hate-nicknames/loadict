package fetch

import (
	"log"
	"sync"
	"time"

	"nvm.ga/loadict/card"
)

const concurrentFetches = 10
const timeout = time.Second * 2

func FetchCards(words []string, fetcher WordFetcher) []*card.Card {
	source := make(chan string, len(words))
	fetched := make(chan *Response, 0)
	rendered := make(chan *card.Card, 0)

	go fetchWords(fetcher, concurrentFetches, source, fetched)
	go renderWords(fetched, rendered)

	cardsEnqueued := 0
	for _, word := range words {
		if cardsEnqueued%concurrentFetches == 0 {
			time.Sleep(timeout)
		}
		source <- word
		cardsEnqueued++
	}
	close(source)

	cards := make([]*card.Card, 0)
	for card := range rendered {
		log.Println("Fetched", card.Word)
		cards = append(cards, card)
	}
	return cards
}

func fetchWords(fetcher WordFetcher, concurrency int, in <-chan string, out chan<- *Response) {
	var wg sync.WaitGroup
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			for word := range in {
				log.Println("Fetching", word)
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
