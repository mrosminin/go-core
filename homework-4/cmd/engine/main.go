package main

import (
	"flag"
	"fmt"
	"go-core-own/homework-4/pkg/index"
	"go-core-own/homework-4/pkg/spider"
	"log"
)

const depth = 2

type Scanner interface {
	Scan(url string) (data map[string]string, err error)
}

type Page struct {
	Url   string
	Title string
}

var urls = []string{
	"https://go.dev",
	"http://www.transflow.ru",
}

func (p Page) String() string {
	return fmt.Sprintf("%s: %s", p.Url, p.Title)
}
func (p Page) Scan(url string) (data map[string]string, err error) {
	return spider.Scan(url, depth)
}

func ScanPages(s Scanner, i *index.Service, pp []string) error {
	for _, p := range pp {
		data, err := s.Scan(p)
		if err != nil {
			return err
		}
		i.Fill(data)
	}
	return nil
}

func main() {
	var str = flag.String("str", "", "Строка для поиска")
	flag.Parse()
	p := new(Page)
	i := index.New()
	err := ScanPages(p, i, urls)
	if err != nil {
		log.Printf("ошибка при сканировании: %v\n", err)
		return
	}
	for {
		for *str == "" {
			fmt.Printf("\nВведите строку для поиска: ")
			fmt.Scanln(str)
		}
		for idx, d := range i.Find(*str) {
			fmt.Printf("%d %v\n", idx+1, Page{Url: d.Url, Title: d.Title})
		}
		*str = ""
	}
}
