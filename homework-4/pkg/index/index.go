package index

import (
	"sort"
	"strings"
)

type Service struct {
	Index map[string][]int
	Docs  []Doc
}

type Doc struct {
	ID    int
	Url   string
	Title string
}

// New - конструктор
func New() *Service {
	var s Service
	s.Index = make(map[string][]int)
	return &s
}

func (s *Service) Fill(data map[string]string) {
	for k, v := range data {
		doc := Doc{ID: len(s.Docs), Url: k, Title: v}
		s.Docs = append(s.Docs, doc)
		ss := strings.Split(v, " ")
		for _, str := range ss {
			str = strings.ToLower(str)
			s.Index[str] = unique(append(s.Index[str], doc.ID))
		}
	}
	// Бессмысленная сортировка по заданию
	sort.Slice(s.Docs, func(i, j int) bool { return s.Docs[i].ID < s.Docs[j].ID })
}

func (s *Service) Find(str string) []Doc {
	var resIdx []int
	for k, v := range s.Index {
		if k == str {
			resIdx = append(resIdx, v...)
		}
	}
	var resDocs []Doc
	for _, idx := range resIdx {
		docIdx := sort.Search(len(s.Docs), func(i int) bool { return s.Docs[i].ID >= idx })
		if docIdx < len(s.Docs) {
			resDocs = append(resDocs, s.Docs[docIdx])
		}
	}
	return resDocs
}

// Возвращает слайс с уникальным набором элементов
func unique(input []int) []int {
	u := make([]int, 0, len(input))
	m := make(map[int]bool)
	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}
