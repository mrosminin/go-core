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
	Url   string
	Title string
}

func (p Page) Scan() (data map[string]string, err error) {
	return spider.Scan(p.Url, depth)
}
func (p Page) String() string {
	return fmt.Sprintf("%s: %s", p.Url, p.Title)
}

type Pages []Page

func (pp Pages) Print() {
	for i, p := range pp {
		fmt.Printf("%d %v\n", i+1, p)
	}
}
func (pp Pages) Find(s string) Pages {
	var res Pages
	for _, p := range pp {
		if strings.Contains(strings.ToLower(p.Title), strings.ToLower(s)) {
			res = append(res, p)
		}
	}
	return res
}

var urls = [2]Page{
	{Url: "https://go.dev"},
	{Url: "http://www.transflow.ru"},
}

func main() {
	var str = flag.String("str", "", "Строка для поиска")
	flag.Parse()
	for *str == "" {
		fmt.Printf("\nВведите строку для поиска: ")
		fmt.Scanln(str)
	}
	for _, p := range urls {
		data, err := Search(p, *str)
		if err != nil {
			log.Printf("ошибка при поиске на странице %s: %v\n", p, err)
			continue
		}
		data.Print()
	}
}

func Search(s Scanner, str string) (Pages, error) {
	var pages Pages
	data, err := s.Scan()
	if err != nil {
		log.Printf("ошибка при сканировании сайта %v\n", err)
		return nil, err
	}
	for k, v := range data {
		pages = append(pages, Page{Url: k, Title: v})
	}
	if len(pages) > 0 {
		pages = pages.Find(str)
	}
	return pages, nil
}
