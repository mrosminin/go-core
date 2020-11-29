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
	gs.storage = storage.New(sl, &btree.Tree{})

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
		for _, s := range gs.sites {
			data, err := gs.scanner.Scan(s, gs.depth)
			if err != nil {
				continue
			}
			gs.engine.Store(data)
		}
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
