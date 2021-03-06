// storage - служба хранения данных

package storage

import (
	"fmt"
	"go-core-own/homework-7/pkg/storage/btree"
)

type SaveLoader interface {
	Load() (p []byte, err error)
	Save(p []byte) error
}

type Service struct {
	SaveLoader
	Docs *btree.Tree
}

func New(sl SaveLoader) *Service {
	return &Service{
		SaveLoader: sl,
		Docs:       &btree.Tree{},
	}
}

// Export экспортирует данные хранилища документов в формате json строки
func (s *Service) Export() {
	jsonData, err := s.Docs.Serialize()
	if err != nil {
		fmt.Printf("Ошибка экспорта: %v", err)
		return
	}
	err = s.Save(jsonData)
	if err != nil {
		fmt.Printf("Ошибка сохранения: %v", err)
		return
	}
}
