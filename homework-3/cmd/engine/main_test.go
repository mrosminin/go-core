package main

import (
	"go-core-own/homework-3/pkg/stub"
	"testing"
)

type StubType int

func (StubType) Scan() (data map[string]string, err error) {
	return stub.Scan()
}

func Test_Scanner(t *testing.T) {
	s := new(StubType)
	want := 2
	got, err := s.Scan()
	if err != nil {
		t.Fatalf("получена ошибка %v", err)
	}
	if len(got) != want {
		t.Errorf("Scan(): получили %d, должны были %d\n", len(got), want)
	}
}
