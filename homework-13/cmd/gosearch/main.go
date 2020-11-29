package main

import (
	"fmt"
	"go-core-own/homework-13/pkg/engine"
	"go-core-own/homework-13/pkg/index"
	"go-core-own/homework-13/pkg/scanner"
	"go-core-own/homework-13/pkg/scanner/spider"
	"go-core-own/homework-13/pkg/storage"
	"go-core-own/homework-13/pkg/storage/btree"
	"go-core-own/homework-13/pkg/storage/diskstor"
	"log"
	"sync"
)

// Поисковик GoSearch
type gosearch struct {
	scanner scanner.Interface
	index   *index.Service
	storage *storage.Service
	engine  *engine.Service

	sites []string
	depth int
}

// Конструктор поисковика
func new() (*gosearch, error) {
	gs := gosearch{
		sites: []string{
			"https://go.dev",
			"http://www.transflow.ru",
			"https://www.newsru.com",
			"https://www.gov-murman.ru/",
			"https://www.anekdot.ru/",
			"https://en.wikipedia.org/wiki/Main_Page",
			"https://www.prj-exp.ru/gost-34",
			"https://habr.com/ru/",
		},
		depth: 2,

		// Определяются зависимости сканер сайтов, служба индексирования
		scanner: spider.New(&spider.Net{}),
		index:   index.New(),
	}

	// Служба хранения данных
	sl, err := diskstor.New("./diskstor.txt")
	if err != nil {
		return nil, err
	}
	gs.storage = storage.New(sl, btree.New())

	// Поисковый движок
	gs.engine = engine.New(gs.index, gs.storage)

	return &gs, nil
}

func main() {
	gs, err := new()
	if err != nil {
		log.Printf("%v\n", err)
		return
	}

	// запуска скинирования страниц, указанных в конструкторе поисковика
	// пока идет сканирование движок выдает результаты из загруженных из долговременного хранилища
	go func() {
		const W = 10                         // кол-во "рабочих"
		sites := make(chan string)           // канал для заданий на сканирование
		res := make(chan []scanner.Document) // канал результатов сканирования
		defer close(res)
		err := make(chan error) // канал с ошибками
		defer close(err)

		var wg sync.WaitGroup
		wg.Add(W)

		for i := 0; i < W; i++ {
			go func(sites <-chan string, results chan<- []scanner.Document, err chan<- error) {
				defer wg.Done()
				for s := range sites {
					data, e := gs.scanner.Scan(s, gs.depth)
					if e != nil {
						err <- e
						continue
					}
					results <- data
				}

			}(sites, res, err)
		}
		// поток для записи результатов сканирования
		go func(results <-chan []scanner.Document) {
			for data := range results {
				gs.engine.Store(data)
			}
		}(res)
		// поток для чтения ошибок
		go func(err <-chan error) {
			for e := range err {
				fmt.Printf("Ошибка при сканировании сайта: %v", e)
			}
		}(err)

		for _, s := range gs.sites {
			sites <- s
		}
		close(sites)
		wg.Wait()
	}()

	var query string
	for {
		for query == "" {
			fmt.Printf("\nВведите строку для поиска: ")
			fmt.Scanln(&query)
		}
		for i, d := range gs.engine.Find(query) {
			fmt.Printf("%d %v\n", i+1, d)
		}
		query = ""
	}
}
