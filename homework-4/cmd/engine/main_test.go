package main

import (
	"go-core-own/homework-4/pkg/index"
	"go-core-own/homework-4/pkg/stub"
	"log"
	"testing"
)

type StubType int

func (StubType) Scan(url string) (data map[string]string, err error) {
	return stub.Scan()
}

func Test_Index(t *testing.T) {
	s := new(StubType)
	err := ScanPages(s, []Page{{Url: "test1"}, {Url: "test2"}})
	if err != nil {
		log.Printf("ошибка при сканировании: %v\n", err)
		return
	}
	want := 2
	got := index.Find("сервис")
	if len(got) != want {
		t.Errorf("index.Find('сервис'): получили %d, должны были %d\n", len(got), want)
	}
}
