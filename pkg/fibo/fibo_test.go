package fibo

import "testing"

func TestFibo(t *testing.T) {
	want := 1
	got := Fibo(2)
	if got != want {
		t.Errorf("Ошибка: получили %d, должны были %d\n", got, want)
	}

	want = 233
	got = Fibo(12)
	if got != want {
		t.Errorf("Ошибка: получили %d, должны были %d\n", got, want)
	}

}
