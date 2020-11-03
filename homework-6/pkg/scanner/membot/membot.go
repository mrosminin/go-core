package membot

import (
	"go-core-own/homework-6/pkg/crawler"
)

// Service - имитация служба поискового робота.
type Service struct{}

// New - констрктор имитации службы поискового робота.
func New() *Service {
	s := Service{}
	return &s
}

// Scan возвращает заранее подготовленный набор данных
func (s *Service) Scan(url string, depth int) ([]scanner.Document, error) {

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

	return data, nil
}
