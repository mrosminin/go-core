package main

import (
	"flag"
	"fmt"
	"go-core-own/homework-5/pkg/engine"
	"log"
)

const depth = 2

var urls = []string{
	"https://go.dev",
	"http://www.transflow.ru",
}

func main() {
	var str = flag.String("str", "", "Строка для поиска")
	flag.Parse()
	e := engine.New()
	for _, p := range urls {
		err := e.ScanPage(p, depth)
		if err != nil {
			log.Printf("ошибка при сканировании: %v\n", err)
			continue
		}
	}
	// поиск документов по строке ввода, либо по строке переданной флагом
	for {
		for *str == "" {
			fmt.Printf("\nВведите строку для поиска: ")
			fmt.Scanln(str)
		}
		for i, d := range e.Find(*str) {
			fmt.Printf("%d %v\n", i+1, d)
		}
		*str = ""
	}
}
