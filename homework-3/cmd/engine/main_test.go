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
	cases := []struct {
		in   string
		want int
	}{
		{"тр", 2},
		{"сервис", 1},
		{"сеновал", 0},
	}
	for _, c := range cases {
		got, err := Search(s, c.in)
		if err != nil {
			t.Errorf("%v\n", err)
		}
		if len(got) != c.want {
			t.Errorf("Search(%s): получили %d, должны были %d\n", c.in, len(got), c.want)
		}
	}
}
