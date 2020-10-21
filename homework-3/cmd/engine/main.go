package main

import (
	"errors"
	"flag"
	"fmt"
	"go-core-own/homework-3/pkg/spider"
	"go-core-own/homework-3/pkg/stub"
	"log"
	"strings"
)

const depth = 2

var sites = [2]myType{"https://go.dev", "http://www.transflow.ru"}

type Page struct {
	Url   string
	Title string
}
type Pages []Page

func (p Page) String() string {
	return fmt.Sprintf("%s: %s", p.Url, p.Title)
}
func (pp Pages) Print() {
	for i, p := range pp {
		fmt.Printf("%d %v\n", i+1, p)
	}
}

func (pp Pages) Find(s string) (Pages, error) {
	var res Pages
	for _, p := range pp {
		if strings.Contains(strings.ToLower(p.Title), strings.ToLower(s)) {
			res = append(res, p)
		}
	}
	if len(res) == 0 {
		return nil, errors.New("ничего не найдено")
	}
	return res, nil
}

type Scanner interface {
	Scan() (data map[string]string, err error)
}

type StubType int

func (st StubType) Scan() (data map[string]string, err error) {
	return stub.Scan()
}

type myType string

func (s myType) Scan() (data map[string]string, err error) {
	return spider.Scan(string(s), depth)
}

func Search(i Scanner, str string) (Pages, error) {
	var pages Pages
	data, err := i.Scan()
	if err != nil {
		log.Printf("ошибка при сканировании сайта %v\n", err)
		return nil, err
	}
	for k, v := range data {
		pages = append(pages, Page{Url: k, Title: v})
	}
	if len(pages) == 0 {
		return nil, errors.New("ссылок не найдено")
	}
	pages, err = pages.Find(str)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return pages, nil
}

func main() {
	var str = flag.String("str", "", "Строка для поиска")
	flag.Parse()
	for *str == "" {
		fmt.Printf("\nВведите строку для поиска: ")
		fmt.Scanln(str)
	}
	for _, s := range sites {
		data, err := Search(s, *str)
		if err != nil {
			log.Printf("ошибка при поиске на сайте %s: %v\n", s, err)
			continue
		}
		data.Print()
	}
}
