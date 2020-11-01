// Engine - служба поискового движка, требует службы сканирования сайтов и индексирования результатов поиска
package engine

import (
	"go-core-own/homework-5/pkg/crawler"
	"go-core-own/homework-5/pkg/crawler/spider"
	"go-core-own/homework-5/pkg/index"
)

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
