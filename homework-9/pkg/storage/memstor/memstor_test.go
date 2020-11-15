package memstor

import (
	"testing"
)

func BenchmarkDiskstor(b *testing.B) {
	s := New(false)
	for i := 0; i < b.N; i++ {
		_ = s.Save([]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor"))
		_, _ = s.Load()
	}
}
