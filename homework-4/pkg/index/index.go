package index

import (
	"go-core-own/homework-4/pkg/crawler"
	"sort"
	"strings"
)

type Indexer interface {
	Fill(data []crawler.Document)
	Find(str string) []crawler.Document
}

type Service struct {
	Index map[string][]int
	Docs  []crawler.Document
}

// New - конструктор
func New() *Service {
	return &Service{
		Index: make(map[string][]int),
	}
}

// Fill - создание обратного индекса
func (s *Service) Fill(data []crawler.Document) {
	for _, d := range data {
		d.ID = len(s.Docs)
		s.Docs = append(s.Docs, d)
		ss := strings.Split(d.Title, " ")
		for _, str := range ss {
			str = strings.ToLower(str)
			if !arrayHasEl(s.Index[str], d.ID) {
				s.Index[str] = append(s.Index[str], d.ID)
			}
		}
	}
	// Бессмысленная сортировка по заданию
	sort.Slice(s.Docs, func(i, j int) bool { return s.Docs[i].ID < s.Docs[j].ID })
}

// Find - поиск страниц по слову в заголоке
func (s *Service) Find(str string) []crawler.Document {
	ids, ok := s.Index[strings.ToLower(str)]
	if !ok {
		return nil
	}
	var result []crawler.Document
	for _, id := range ids {
		idx := sort.Search(len(s.Docs), func(i int) bool { return s.Docs[i].ID >= id })
		if idx < len(s.Docs) {
			result = append(result, s.Docs[idx])
		}
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
