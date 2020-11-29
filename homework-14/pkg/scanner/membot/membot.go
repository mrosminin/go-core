// membot реализует имитацию поискового робота
package membot

import (
	"go-core-own/homework-14/pkg/scanner"
)

type Service struct{}

// Scan возвращает заранее подготовленный набор данных
func (s *Service) Scan(string, int) ([]scanner.Document, error) {
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
