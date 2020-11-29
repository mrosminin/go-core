// engine - служба поискового движка, требует службы индексирования результатов поиска и хранения
// Функции службы:
// 1. Выдача документов по строке запроса (берет id в службе индексирования, а затем выдает необходимые документы из хранилища)
// 2. Импорт/экспорт предыдущих результатов сканирования
package engine

import (
	"go-core-own/homework-14/pkg/index"
	"go-core-own/homework-14/pkg/scanner"
	"go-core-own/homework-14/pkg/storage/btree"
)

type Service struct {
	index   *index.Service
	storage *btree.Tree
}

func New(index *index.Service, storage *btree.Tree) *Service {
	return &Service{
		index:   index,
		storage: storage,
	}
}

// Store сохраняет документы в хранилище, передает на индексирование, экспортирует в долгосрочное хранилище
func (s *Service) Store(data []scanner.Document) error {
	for _, d := range data {
		d.ID = s.storage.Insert(d)
		s.index.Add(d)
	}
	return nil
}

// Find выдает результаты поискового запроса (id документов берутся из службы индексирования, сами документы из хранилища)
func (s *Service) Find(q string) []scanner.Document {
	ids := s.index.Find(q)
	var result []scanner.Document
	for _, id := range ids {
		doc, err := s.storage.Find(id)
		if err != nil {
			continue
		}
		result = append(result, doc)
	}
	return result
}
