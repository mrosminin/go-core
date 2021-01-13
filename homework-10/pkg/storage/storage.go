// slice - пакет для хранения документов в массиве
package storage

import (
	"errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go-core-own/homework-10/pkg/scanner"
)

var (
	storageDocumentsNumber = promauto.NewCounter(prometheus.CounterOpts{
		Name: "storage_documents_number",
		Help: "Количество документов в хранилище.",
	})
)

type Service struct {
	Docs []scanner.Document
}

// Insert вставляет элемент в массив
func (s *Service) Insert(doc scanner.Document) (id int) {
	doc.ID = len(s.Docs)
	s.Docs = append(s.Docs, doc)
	storageDocumentsNumber.Inc()
	return doc.ID
}

// Find находит элемент в массиве по ID документа
func (s *Service) Find(id int) (scanner.Document, error) {
	for _, doc := range s.Docs {
		if doc.ID == id {
			return doc, nil
		}
	}
	return scanner.Document{}, errors.New("документ не найден")
}

func (s *Service) Update(doc scanner.Document) error {
	for i := 0; i < len(s.Docs); i++ {
		if s.Docs[i].ID == doc.ID {
			s.Docs[i] = doc
			return nil
		}
	}
	return errors.New("документ не найден")
}

func (s *Service) Delete(doc scanner.Document) error {
	var docs []scanner.Document
	for _, doc := range s.Docs {
		if doc.ID != doc.ID {
			docs = append(docs, doc)
		}
	}
	if len(docs) != len(s.Docs) {
		s.Docs = docs
		return nil
	}
	return errors.New("документ не найден")
}
