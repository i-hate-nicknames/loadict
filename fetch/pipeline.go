package fetch

import (
	"log"
	"sync"
)

const concurrentFetches = 10

// todo: return cards, get rid of ExportCard

func FetchCards(words []string, fetcher WordFetcher) <-chan *ExportCard {
	source := make(chan string, len(words))
	fetched := make(chan *Response, 0)
	rendered := make(chan *ExportCard, 0)

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

func renderWords(in <-chan *Response, out chan<- *ExportCard) {
	for wordResponse := range in {
		log.Println("Rendering", wordResponse.Word)
		card, err := renderCard(wordResponse)
		if err != nil {
			log.Println(err)
			continue
		}
		out <- &ExportCard{Word: wordResponse.Word, Card: card}
	}
	close(out)
}

type ExportCard struct {
	Word, Card string
}
