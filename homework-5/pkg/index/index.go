package index

import (
	"fmt"
	"go-core-own/homework-5/pkg/crawler"
	"go-core-own/homework-5/pkg/index/btree"
	"strings"
)

type Indexer interface {
	Fill(data []crawler.Document)
	Find(str string) []crawler.Document
}

type Service struct {
	Index map[string][]int
	Docs  *btree.Tree
}

// New - конструктор
func New() *Service {
	return &Service{
		Index: make(map[string][]int),
		Docs:  &btree.Tree{},
	}
}

// Fill - создание обратного индекса
func (s *Service) Fill(data []crawler.Document) {
	for _, d := range data {
		id := s.Docs.Insert(d)
		fmt.Println(id, d.Title)
		ss := strings.Split(d.Title, " ")
		for _, str := range ss {
			str = strings.ToLower(str)
			if !arrayHasEl(s.Index[str], id) {
				s.Index[str] = append(s.Index[str], id)
			}
		}
	}
}

// Find - поиск страниц по слову в заголоке
func (s *Service) Find(str string) []crawler.Document {
	ids, ok := s.Index[strings.ToLower(str)]
	if !ok {
		return nil
	}
	var result []crawler.Document
	for _, id := range ids {
		doc, err := s.Docs.Find(id)
		if err != nil {
			continue
		}
		result = append(result, doc)
	}
	return result
}

// Проверка наличия в массиве элемента
func arrayHasEl(slice []int, el int) bool {
	for _, val := range slice {
		if val == el {
			return true
		}
	}
	return false
}
