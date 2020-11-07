// index - служба индексирования
// Задача - вести ассоциативный массив слово: массив id документов в хранилище, в которых встречается это слово
package index

import (
	"go-core-own/homework-7/pkg/scanner"
	"strings"
)

type Service struct {
	Index map[string][]int
}

func New() *Service {
	return &Service{
		Index: make(map[string][]int),
	}
}

// Insert добавляет в индекс отдельные слова из заголовка документа. У документа уже есть ID
func (s *Service) Insert(d scanner.Document) {
	ss := strings.Split(d.Title, " ")
	for _, str := range ss {
		str = strings.ToLower(str)
		if !contains(s.Index[str], d.ID) {
			s.Index[str] = append(s.Index[str], d.ID)
		}
	}
}

// Find возвращает массив id документов, в которых было найдено слово
func (s *Service) Find(q string) []int {
	return s.Index[strings.ToLower(q)]
}

// contains - проверка наличия в массиве элемента
func contains(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
