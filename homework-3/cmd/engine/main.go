package main

import (
	"flag"
	"fmt"
	"go-core-own/homework-2/pkg/spider"
	"go-core-own/homework-2/pkg/stub"
	"log"
	"strings"
)

const depth = 2

var sites = [2]str{"https://go.dev", "http://www.transflow.ru"}

type Page struct {
	Url   string
	Title string
}
type Pages []Page

func (p Page) String() string {
	return fmt.Sprintf("%s: %s", p.Url, p.Title)
}
func (pp Pages) Print() {
	if len(pp) == 0 {
		fmt.Println("Ничего не найдено. Попробуй еще.")
		return
	}
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

type Scanner interface {
	Scan()
}

type stubType int

func (st stubType) Scan() (data map[string]string, err error) {
	return stub.Scan()
}

type str string

func (s str) Scan() (data map[string]string, err error) {
	return spider.Scan(string(s), depth)
}

func main() {
	var pages Pages
	for _, s := range sites {
		data, err := s.Scan()
		if err != nil {
			log.Printf("ошибка при сканировании сайта %s: %v\n", s, err)
			continue
		}
		for k, v := range data {
			pages = append(pages, Page{Url: k, Title: v})
		}
	}
	fmt.Printf("На сайтах %v найдено %d ссылок\n", sites, len(pages))

	var str = flag.String("str", "", "Строка для поиска")
	flag.Parse()

	if *str != "" {
		pages.Find(*str).Print()
		return
	}
	for {
		fmt.Printf("\nВведите строку для поиска: ")
		fmt.Scanln(str)
		pages.Find(*str).Print()
	}
}
