package fibo

import "testing"

func TestFibo(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{7, 21},
		{8, 34},
	}
	for _, c := range cases {
		got := Fibo(c.in)
		if got != c.want {
			t.Errorf("Fibo(%d): получили %d, должны были %d\n", c.in, got, c.want)
		}
	}
}
