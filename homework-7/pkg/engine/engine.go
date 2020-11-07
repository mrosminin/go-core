// engine - служба поискового движка, требует службы индексирования результатов поиска и хранения
// Функции службы:
// 1. Выдача документов по строке запроса (берет id в службе индексирования, а затем выдает необходимые документы из хранилища)
// 2. Импорт предыдущих результатов сканирования
package engine

import (
	"encoding/json"
	"go-core-own/homework-7/pkg/index"
	"go-core-own/homework-7/pkg/scanner"
	"go-core-own/homework-7/pkg/storage"
	"log"
)

type Service struct {
	index   *index.Service
	storage *storage.Service
}

func New(index *index.Service, storage *storage.Service) *Service {
	s := &Service{
		index:   index,
		storage: storage,
	}
	s.init()
	return s
}

// Инициализация поискового движка. Загрузка данных прошлых сканирований в индекс и хранилище документов
func (s *Service) init() {
	jsonData, err := s.storage.Load()
	if err != nil {
		log.Printf("ошибка чтения из хранилища: %v\n", err)
		return
	}
	s.Import(jsonData)
}

// Store сохраняет документы в хранилище, передает на индексирование, экспортирует в долгосрочное хранилище
func (s *Service) Store(data []scanner.Document) {
	for _, d := range data {
		d.ID = s.storage.Docs.Insert(d)
		s.index.Insert(d)
	}
	s.storage.Export()
}

// Find выдает результаты поискового запроса (id документов берутся из службы индексирования, сами документы из хранилища)
func (s *Service) Find(q string) []scanner.Document {
	ids := s.index.Find(q)
	var result []scanner.Document
	for _, id := range ids {
		doc, err := s.storage.Docs.Find(id)
		if err != nil {
			continue
		}
		result = append(result, doc)
	}
	return result
}

// Import декодирует данные из json строки, передает на хранение
func (s *Service) Import(p []byte) error {
	var docs []scanner.Document
	err := json.Unmarshal(p, &docs)
	if err != nil {
		return err
	}
	go s.Store(docs)
	return nil
}
