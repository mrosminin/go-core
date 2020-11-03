package scanner

import (
	"fmt"
	"log"
)

// Document - документ, веб-страница, полученная поисковым роботом.
type Document struct {
	ID    int
	URL   string
	Title string
	Body  string
}

func (d Document) String() string {
	return fmt.Sprintf("%s: %s", d.URL, d.Title)
}

// Интерфейс поискового робота.
type Interface interface {
	Scan(url string, depth int) ([]Document, error)
}

// Cлужба скинирования сайтов
type Service struct {
	Interface
}

func New(sc Interface) *Service {
	return &Service{sc}
}

// ScanPages - метод сканированияя страниц
func (s *Service) ScanPages(urls []string, depth int) (data []Document) {
	for _, p := range urls {
		dd, err := s.Scan(p, depth)
		if err != nil {
			log.Printf("ошибка при сканировании: %v\n", err)
			continue
		}
		data = append(data, dd...)
	}
	return data
}
