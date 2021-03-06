package main

import (
	"flag"
	"fmt"
	"go-core-own/homework-4/pkg/crawler"
	"go-core-own/homework-4/pkg/crawler/spider"
	"go-core-own/homework-4/pkg/index"
	"log"
)

const depth = 2

var urls = []string{
	"https://go.dev",
	"http://www.transflow.ru",
}

// Engine - cлужба поискового движка, требует службы сканирования сайтов и индексирования результатов поиска
type Engine struct {
	crawler crawler.Scanner
	index   index.Indexer
}

// New - Конструктор службы, создает экземпляры служб сканирования и индексирования
func New() *Engine {
	return &Engine{
		crawler: spider.New(),
		index:   index.New(),
	}
}

// ScanPage - метод для сканированияя страницы. Сканирует, индексирут
func (e *Engine) ScanPage(url string, depth int) error {
	data, err := e.crawler.Scan(url, depth)
	if err != nil {
		return err
	}
	e.index.Fill(data)
	return nil
}

// Find - метод для поиска слова в результатах сканирования
func (e *Engine) Find(str string) []crawler.Document {
	return e.index.Find(str)
}

func main() {
	var str = flag.String("str", "", "Строка для поиска")
	flag.Parse()
	e := New()
	for _, p := range urls {
		err := e.ScanPage(p, depth)
		if err != nil {
			log.Printf("ошибка при сканировании: %v\n", err)
			continue
		}
	}
	// поиск документов по строке ввода, либо по строке переданной флагом
	for {
		for *str == "" {
			fmt.Printf("\nВведите строку для поиска: ")
			fmt.Scanln(str)
		}
		for i, d := range e.Find(*str) {
			fmt.Printf("%d %v\n", i+1, d)
		}
		*str = ""
	}
}
