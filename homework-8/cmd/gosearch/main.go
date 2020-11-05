package main

import (
	"fmt"
	"go-core-own/homework-8/pkg/engine"
	"go-core-own/homework-8/pkg/index"
	"go-core-own/homework-8/pkg/scanner"
	"go-core-own/homework-8/pkg/scanner/spider"
	"go-core-own/homework-8/pkg/storage"
	"go-core-own/homework-8/pkg/storage/diskstor"
	"log"
)

// Поисковик GOSearch
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
	// Определяются зависимости поисковый движок, сканер сайтов, служба индексирования, служба хранения
	idx := index.New()
	sc := spider.New()
	sl, err := diskstor.New()
	if err != nil {
		return nil, err
	}
	st := storage.New(sl)
	en := engine.New(idx, st)

	return &gosearch{
		scanner: sc,
		storage: st,
		engine:  en,
		index:   idx,

		sites: []string{
			"https://go.dev",
			"http://www.transflow.ru",
			"https://www.newsru.com",
			"https://www.gov-murman.ru/",
		},
		depth: 2,
	}, nil
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
		ch := make(chan []scanner.Document, len(gs.sites))
		for _, s := range gs.sites {
			go gs.scanner.Scan(s, gs.depth, ch)
			go gs.engine.Store(<-ch)
		}
		close(ch)
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
