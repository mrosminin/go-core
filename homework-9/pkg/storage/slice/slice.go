// slice - пакет для хранения документов в массиве
package slice

import (
	"errors"
	"go-core-own/homework-9/pkg/scanner"
	"go-core-own/homework-9/pkg/storage"
)

type Slice struct {
	Docs []scanner.Document
	storage.Interface
}

// Insert вставляет элемент в массив
func (s *Slice) Insert(doc scanner.Document) (id int) {
	doc.ID = len(s.Docs)
	s.Docs = append(s.Docs, doc)
	return doc.ID
}

// Find находит элемент в массиве по ID документа
// Ради бенчмарка сделаем вид, что id документа может не соответствовать индексу в массиве
func (s *Slice) Find(id int) (scanner.Document, error) {
	for _, doc := range s.Docs {
		if doc.ID == id {
			return doc, nil
		}
	}
	return scanner.Document{}, errors.New("документ не найден")
}
