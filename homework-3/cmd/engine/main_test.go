package main

import (
	"testing"
)

func Test_Scanner(t *testing.T) {
	st := new(StubType)
	cases := []struct {
		in   string
		want int
	}{
		{"тр", 1},
	}
	for _, c := range cases {
		got, err := Search(st, c.in)
		if err != nil {
			t.Errorf("%v\n", err)
		}
		if len(got) != c.want {
			t.Errorf("Search(%s): получили %d, должны были %d\n", c.in, len(got), c.want)
		}
	}
}
