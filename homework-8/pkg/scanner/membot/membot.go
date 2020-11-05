// membot реализует имитацию поискового робота
package membot

import (
	"go-core-own/homework-8/pkg/scanner"
)

type Service struct{}

func New() *Service {
	s := Service{}
	return &s
}

// Scan возвращает заранее подготовленный набор данных
func (s *Service) Scan(url string, depth int, ch chan<- []scanner.Document) {

	data := []scanner.Document{
		{
			ID:    0,
			URL:   "https://yandex.ru",
			Title: "Яндекс",
		},
		{
			ID:    1,
			URL:   "https://google.ru",
			Title: "Google",
		},
	}

	ch <- data
}
