// storage - служба хранения данных

package storage

import (
	"go-core-own/homework-8/pkg/storage/btree"
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
