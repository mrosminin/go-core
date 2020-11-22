package slice

import (
	"go-core-own/homework-9/pkg/scanner"
	"testing"
)

func BenchmarkSlice_InsertFind(b *testing.B) {
	s := Slice{}
	for i := 0; i < b.N; i++ {
		id := s.Insert(scanner.Document{
			URL:   "htt://url.com",
			Title: "Title",
		})
		_, _ = s.Find(id)
	}
}
