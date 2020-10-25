package index

import (
	"sort"
	"strings"
)

// Обратный индекс - хранит по ключу (слову) массив id, документов где данное слово встречается
var Index = make(map[string][]int)

// Документы - массив элементов структуры с id, на которые ссылается индекс
type Doc struct {
	ID    int
	Url   string
	Title string
}

var Docs []Doc

func Fill(data map[string]string) {
	for k, v := range data {
		doc := Doc{ID: len(Docs), Url: k, Title: v}
		Docs = append(Docs, doc)
		ss := strings.Split(v, " ")
		for _, s := range ss {
			s = strings.ToLower(s)
			Index[s] = append(Index[s], doc.ID)
		}
	}
	sort.Slice(Docs, func(i, j int) bool { return Docs[i].ID < Docs[j].ID })
}

func Find(s string) []Doc {
	var resIdx []int
	for k, v := range Index {
		if strings.Contains(k, s) {
			resIdx = append(resIdx, v...)
		}
	}
	var resDocs []Doc
	for _, idx := range unique(resIdx) {
		docIdx := sort.Search(len(Docs), func(i int) bool { return Docs[i].ID >= idx })
		if docIdx < len(Docs) {
			resDocs = append(resDocs, Docs[docIdx])
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
