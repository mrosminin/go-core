// engine - служба поискового движка, требует службы индексирования результатов поиска и хранения
// Функции службы:
// 1. Выдача документов по строке запроса (берет id в службе индексирования, а затем выдает необходимые документы из хранилища)
// 2. Импорт/экспорт предыдущих результатов сканирования
package engine

import (
	"go-core-own/homework-10/pkg/index"
	"go-core-own/homework-10/pkg/scanner"
	"go-core-own/homework-10/pkg/storage"
)

type Service struct {
	index   *index.Service
	Storage *storage.Service
}

func New(index *index.Service, storage *storage.Service) *Service {
	return &Service{
		index:   index,
		Storage: storage,
	}
}

// Store сохраняет документы в хранилище, передает на индексирование, экспортирует в долгосрочное хранилище
func (s *Service) Store(data []scanner.Document) {
	for _, d := range data {
		d.ID = s.Storage.Insert(d)
		s.index.Add(d)
	}
}

// Find выдает результаты поискового запроса (id документов берутся из службы индексирования, сами документы из хранилища)
func (s *Service) Find(q string) []scanner.Document {
	ids := s.index.Find(q)
	var result []scanner.Document
	for _, id := range ids {
		doc, err := s.Storage.Find(id)
		if err != nil {
			continue
		}
		result = append(result, doc)
	}
	return result
}
