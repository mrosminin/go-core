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
	for _, p := range urls {
		data, err := p.Scan()
		if err != nil {
			log.Printf("ошибка при поиске на странице %s: %v\n", p, err)
			continue
		}
		index.Fill(data)
	}
	for {
		for *str == "" {
			fmt.Printf("\nВведите строку для поиска: ")
			fmt.Scanln(str)
		}
		for i, d := range index.Find(*str) {
			fmt.Printf("%d %v\n", i+1, Page{Url: d.Url, Title: d.Title})
		}
		*str = ""
	}
}
