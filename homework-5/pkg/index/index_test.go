package index

import (
	"go-core-own/homework-5/pkg/crawler"
	"testing"
)

func Test_Find(t *testing.T) {
	idx := New()
	fixtures := []crawler.Document{
		{URL: "url1", Title: "заголовок из нескольких слов"},
		{URL: "url2", Title: "иЗ нЕсКоЛЬких сЛОв"},
		{URL: "url3", Title: "_нескольких_слов_"},
	}
	idx.Fill(fixtures)

	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Тест №1",
			s:    "заголовок",
			want: 1,
		},
		{
			name: "Тест №2",
			s:    "нЕсКоЛьКИх",
			want: 2,
		},
		{
			name: "Тест №3",
			s:    "абракадабра",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := idx.Find(tt.s); len(got) != tt.want {
				t.Errorf("Find('%s') = %v, want %v", tt.s, len(got), tt.want)
			}
		})
	}
}
