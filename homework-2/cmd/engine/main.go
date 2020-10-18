package main

import (
	"flag"
	"fmt"
	"go-core-own/homework-2/pkg/spider"
	"log"
	"strings"
)

type Page struct {
	Url   string
	Title string
}

func (p Page) String() string {
	return fmt.Sprintf("%s: %s", p.Url, p.Title)
}

func main() {
	var url = flag.String("url", "", "URL для сканирования")
	var str = flag.String("str", "", "Строка для поиска")
	flag.Parse()

	for i := 0; i < 2; i++ {
		// Разумно полагаю, что если урл задан ключом, то следует просканировать только одну страницу
		if *url != "" {
			i++
		}
		for *url == "" {
			fmt.Printf("Введите URL для сканирования: ")
			fmt.Scanln(url)
		}

		data, err := spider.Scan(*url, 2)
		if err != nil {
			log.Printf("ошибка при сканировании сайта %s: %v\n", *url, err)
			*url = ""
			continue
		}
		var pages []Page
		for k, v := range data {
			pages = append(pages, Page{Url: k, Title: v})
		}
		if len(pages) == 0 {
			fmt.Println("На странице не найдено ни одной ссылки.")
		}
		if len(pages) > 0 {
			if *str == "" {
				fmt.Printf("На странице %s найдено %d ссылок\n", *url, len(pages))
			}
			for *str == "" {
				fmt.Printf("Введите строку для поиска: ")
				fmt.Scanln(str)
			}
			var res []Page
			for _, p := range pages {
				if strings.Contains(strings.ToLower(p.Title), strings.ToLower(*str)) {
					res = append(res, p)
				}
			}
			if len(res) == 0 {
				fmt.Println("Яндекс - найдётся всё! А у нас краулер.")
			}
			if len(res) > 0 {
				fmt.Printf("Найдено %d совпадений:\n", len(res))
				for i, p := range res {
					fmt.Printf("%d. %v\n", i+1, p)
				}
			}
		}
		if i == 0 {
			*url = ""
			*str = ""
			fmt.Println("Требование такое - пару раз просканировать.")
		}
	}

}
