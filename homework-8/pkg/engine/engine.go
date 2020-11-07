// engine - служба поискового движка, требует службы индексирования результатов поиска и хранения
// Функции службы:
// 1. Выдача документов по строке запроса (берет id в службе индексирования, а затем выдает необходимые документы из хранилища)
// 2. Импорт предыдущих результатов сканирования
package engine

import (
	"encoding/json"
	"fmt"
	"go-core-own/homework-8/pkg/index"
	"go-core-own/homework-8/pkg/scanner"
	"go-core-own/homework-8/pkg/storage"
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
	err := s.init()
	if err != nil {
		// Ошибку по стеку не возвращаем, т.к. система и без загруженных данных будет функционировать
		fmt.Println(err)
	}
	return s
}

// Инициализация поискового движка. Загрузка данных прошлых сканирований в индекс и хранилище документов
func (s *Service) init() error {
	jsonData, err := s.storage.Load()
	if err != nil {
		return err
	}
	err = s.restore(jsonData)
	if err != nil {
		return err
	}
	return nil
}

// Store сохраняет документы в хранилище, передает на индексирование, экспортирует в долгосрочное хранилище
func (s *Service) Store(data []scanner.Document) error {
	for _, d := range data {
		d.ID = s.storage.Docs.Insert(d)
		s.index.Add(d)
	}
	err := s.export()
	if err != nil {
		return err
	}
	return nil
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

// restore загружает данные прошлых сканирований из долговременного хранилища
// декодирует данные из json строки, передает на хранение
func (s *Service) restore(p []byte) error {
	var docs []scanner.Document
	err := json.Unmarshal(p, &docs)
	if err != nil {
		return err
	}
	err = s.Store(docs)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) export() error {
	jsonData, err := s.storage.Docs.Serialize()
	if err != nil {
		return err
	}
	err = s.storage.Save(jsonData)
	if err != nil {
		return err
	}
	return nil
}
