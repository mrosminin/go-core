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
	Scan() (data map[string]string, err error)
}

type Page struct {
	Scanner
	Url   string
	Title string
}

func (p Page) Scan() (data map[string]string, err error) {
	return spider.Scan(p.Url, depth)
}
func (p Page) String() string {
	return fmt.Sprintf("%s: %s", p.Url, p.Title)
}

var urls = [2]Page{
	{Url: "https://go.dev"},
	{Url: "http://www.transflow.ru"},
}

func main() {
	var str = flag.String("str", "", "Строка для поиска")
	flag.Parse()

	// Сканируем страницы, сохраняем результаты сканирования с pages
	var pages []Page
	for _, p := range urls {
		data, err := p.Scan()
		if err != nil {
			log.Printf("ошибка при поиске на странице %s: %v\n", p, err)
			continue
		}
		for k, v := range data {
			pages = append(pages, Page{Url: k, Title: v})
		}
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
