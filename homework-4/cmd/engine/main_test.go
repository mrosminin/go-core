package main

import (
	"go-core-own/homework-4/pkg/index"
	"go-core-own/homework-4/pkg/stub"
	"testing"
)

type StubType int

func (st StubType) Scan(url string) (data map[string]string, err error) {
	return stub.Scan()
}

func Test_Scanner(t *testing.T) {
	stub := new(StubType)
	i := index.New()
	err := ScanPages(stub, i, []string{"http://..."})
	if err != nil {
		t.Fatalf("получена ошибка %v", err)
	}
	want := 2
	got := i.Find("трансфлоу")
	if len(got) != want {
		t.Errorf("Scan(): получили %d, должны были %d\n", len(got), want)
	}
}
