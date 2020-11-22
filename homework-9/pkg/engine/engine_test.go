package engine

import (
	"go-core-own/homework-9/pkg/index"
	"go-core-own/homework-9/pkg/scanner"
	"go-core-own/homework-9/pkg/storage"
	"go-core-own/homework-9/pkg/storage/btree"
	"go-core-own/homework-9/pkg/storage/memstor"
	"os"
	"reflect"
	"testing"
)

var s *Service
var fixtures []scanner.Document

func TestMain(m *testing.M) {
	s = New(index.New(), storage.New(memstor.New(false), &btree.Tree{}))
	fixtures = []scanner.Document{
		{ID: 0, URL: "url1", Title: "заголовок из нескольких слов заголовок из нескольких слов"},
		{ID: 1, URL: "url2", Title: "ЗаГолоВОК иЗ нЕсКоЛЬких сЛОв"},
		{ID: 2, URL: "url3", Title: "другой заголовок"},
	}
	s.Store(fixtures)
	os.Exit(m.Run())
}

func TestService_Find(t *testing.T) {
	tests := []struct {
		name  string
		query string
		want  []scanner.Document
	}{
		{
			name:  "Тест1",
			query: "заголовок",
			want:  fixtures,
		},
		{
			name:  "Тест2",
			query: "НЕСКОЛЬких",
			want:  []scanner.Document{fixtures[0], fixtures[1]},
		},
		{
			name:  "Тест3",
			query: "другой",
			want:  []scanner.Document{fixtures[2]},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.Find(tt.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
