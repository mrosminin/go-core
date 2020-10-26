package main

import (
	"go-core-own/homework-3/pkg/stub"
	"testing"
)

type StubType int

func (StubType) Scan(url string) (data map[string]string, err error) {
	return stub.Scan()
}

func Test_Scanner(t *testing.T) {
	s := new(StubType)
	want := 4
	got, err := ScanPages(s, []Page{{Url: "test1"}, {Url: "test2"}})
	if err != nil {
		t.Fatalf("получена ошибка %v", err)
	}
	if len(got) != want {
		t.Errorf("Scan(): получили %d, должны были %d\n", len(got), want)
	}
}
