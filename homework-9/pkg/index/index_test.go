package index

import (
	"go-core-own/homework-9/pkg/scanner"
	"os"
	"reflect"
	"testing"
)

var s *Service
var fixtures []scanner.Document

func TestMain(m *testing.M) {
	s = New()
	fixtures = []scanner.Document{
		{ID: 1, URL: "url1", Title: "заголовок из нескольких слов заголовок из нескольких слов"},
		{ID: 2, URL: "url2", Title: "ЗаГолоВОК иЗ нЕсКоЛЬких сЛОв"},
		{ID: 3, URL: "url3", Title: "другой заголовок"},
	}
	os.Exit(m.Run())
}

func TestService_Add(t *testing.T) {
	for _, d := range fixtures {
		s.Add(d)
	}

	want := map[string][]int{
		"заголовок":  {1, 2, 3},
		"из":         {1, 2},
		"нескольких": {1, 2},
		"слов":       {1, 2},
		"другой":     {3},
	}
	if !reflect.DeepEqual(s.Index, want) {
		t.Errorf("Add() = %v, want %v", s.Index, want)
	}
}

func TestService_Find(t *testing.T) {
	for _, d := range fixtures {
		s.Add(d)
	}
	tests := []struct {
		name string
		q    string
		want []int
	}{
		{
			name: "Тест1",
			q:    "заголовок",
			want: []int{1, 2, 3},
		},
		{
			name: "Тест2",
			q:    "ЗаГоЛовоК",
			want: []int{1, 2, 3},
		},
		{
			name: "Тест2",
			q:    "друГОЙ",
			want: []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.Find(tt.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
