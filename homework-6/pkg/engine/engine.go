// Engine - служба поискового движка, требует службы индексирования результатов поиска и хранения
// Функции движка:
// 1. Выдача документов по строке запроса (ищет индексы в службе индексирования, а затем выдает необходимые документы из хранилища)
// 2. Загрузка документов из долговременного хранилища и передача их службам индексирования и хранения документов
package engine

import (
	"encoding/json"
	"go-core-own/homework-6/pkg/index"
	"go-core-own/homework-6/pkg/scanner"
	"go-core-own/homework-6/pkg/storage"
	"log"
)

// Service - служба поискового движка
type Service struct {
	index   *index.Service
	storage *storage.Service
}

// New - конструктор поискового движка
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
	go s.Import(jsonData)
}

// Store - сохранение найденных документов (в индексе, в хранилище документов в памяти - дереве и в долгосрочном хранилище)
func (s *Service) Store(data []scanner.Document) {
	for _, d := range data {
		d.ID = s.storage.Docs.Insert(d)
		s.index.Insert(d)
	}
	s.storage.Export()
}

// Find - метод выдачи результатов поискового запроса (id документов из индекса, сами документы из хранилища)
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

// Import - импорт данных из json строки
func (s *Service) Import(p []byte) error {
	var docs []scanner.Document
	err := json.Unmarshal(p, &docs)
	if err != nil {
		return err
	}
	go s.Store(docs)
	return nil
}
