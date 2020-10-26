package main

import (
	"flag"
	"fmt"
	"go-core-own/homework-3/pkg/spider"
	"log"
	"strings"
)

const depth = 2

type Scanner interface {
	Scan(url string) (data map[string]string, err error)
}

type Page struct {
	Url   string
	Title string
}

func (p Page) String() string {
	return fmt.Sprintf("%s: %s", p.Url, p.Title)
}

var urls = []Page{
	{Url: "https://go.dev"},
	{Url: "http://www.transflow.ru"},
}

func (p Page) Scan(url string) (data map[string]string, err error) {
	return spider.Scan(url, depth)
}

func ScanPages(s Scanner, pp []Page) ([]Page, error) {
	var pages []Page
	for _, p := range pp {
		data, err := s.Scan(p.Url)
		if err != nil {
			return nil, err
		}
		for k, v := range data {
			pages = append(pages, Page{Url: k, Title: v})
		}
	}
	return pages, nil
}

func main() {
	var str = flag.String("str", "", "Строка для поиска")
	flag.Parse()
	s := new(Page)
	pages, err := ScanPages(s, urls)
	if err != nil {
		log.Printf("ошибка при сканировании: %v\n", err)
		return
	}
	for {
		for *str == "" {
			fmt.Printf("\nВведите строку для поиска: ")
			fmt.Scanln(str)
		}
		i := 1
		for _, p := range pages {
			if strings.Contains(strings.ToLower(p.Title), strings.ToLower(*str)) {
				fmt.Printf("%d %v\n", i, p)
				i++
			}
		}
		*str = ""
	}
}
