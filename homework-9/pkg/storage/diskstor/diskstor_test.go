package diskstor

import (
	"log"
	"testing"
)

func BenchmarkDiskstor_SaveLoad(b *testing.B) {
	s, err := New("benchmark.txt")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		_ = s.Save([]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor"))
		_, _ = s.Load()
	}
}
