// slice - пакет для хранения документов в массиве
package storage

import (
	"errors"
	"go-core-own/homework-19/pkg/scanner"
)

type Service struct {
	Docs []scanner.Document
}

// Insert вставляет элемент в массив
func (s *Service) Insert(doc scanner.Document) (id int) {
	doc.ID = len(s.Docs)
	s.Docs = append(s.Docs, doc)
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
