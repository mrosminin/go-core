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
		ss := strings.Split(d.Title, " ")
		for _, str := range ss {
			str = strings.ToLower(str)
			if !arrayHasEl(s.Index[str], id) {
				s.Index[str] = append(s.Index[str], id)
			}
		}
		s.save()
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

func (s *Service) save() error {
	f, err := os.Create("./file.txt")
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(s.Docs)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(f.Name(), []byte(jsonData), 0666)
	if err != nil {
		log.Fatal(err)
	}
	/*data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Данные файла:\n%s\n", data)
	// вариант используя io.Reader
	file, err := os.Open(f.Name())
	if err != nil {
		log.Fatal(err)
	}*/
	return nil
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
