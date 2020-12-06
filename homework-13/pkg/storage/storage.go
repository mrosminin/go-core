// storage - служба хранения данных

package storage

import (
	"go-core-own/homework-13/pkg/scanner"
)

type SaveLoader interface {
	Load() (p []byte, err error)
	Save(p []byte) error
}

type Interface interface {
	Insert(doc scanner.Document) (id int)
	Find(id int) (scanner.Document, error)
	Json() (jsonData []byte, err error)
}

type Service struct {
	SaveLoader
	Interface
}

func New(sl SaveLoader, m Interface) *Service {
	return &Service{
		sl,
		m,
	}
}
