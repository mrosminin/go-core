package index

import (
	"go-core-own/homework-6/pkg/scanner"
	"strings"
)

type Interface interface {
	Insert(data []scanner.Document)
	Find(str string) []scanner.Document
}

// Service - служба индексирования
type Service struct {
	Index map[string][]int
}

// New - конструктор
func New() *Service {
	return &Service{
		Index: make(map[string][]int),
	}
}

// Insert - создание обратного индекса для документа. У документа уже есть ID
func (s *Service) Insert(d scanner.Document) {
	ss := strings.Split(d.Title, " ")
	for _, str := range ss {
		str = strings.ToLower(str)
		if !arrayHasEl(s.Index[str], d.ID) {
			s.Index[str] = append(s.Index[str], d.ID)
		}
	}
}

// Find - поиск страниц по слову в заголоке
func (s *Service) Find(q string) []int {
	return s.Index[strings.ToLower(q)]
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
