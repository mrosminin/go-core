package btree

import (
	"go-core-own/homework-14/pkg/scanner"
	"os"
	"reflect"
	"testing"
)

var tree *Tree
var fixtures []scanner.Document

func TestMain(m *testing.M) {
	fixtures = []scanner.Document{
		{URL: "url1", Title: "заголовок из нескольких слов"},
		{URL: "url2", Title: "иЗ нЕсКоЛЬких сЛОв"},
		{URL: "url3", Title: "другой заголовок"},
	}
	tree = New()
	os.Exit(m.Run())
}

func TestTree_Insert(t1 *testing.T) {
	tests := []struct {
		name string
		doc  scanner.Document
		want int
	}{
		{
			name: "Тест 1",
			doc:  fixtures[0],
			want: 0,
		},
		{
			name: "Тест 2",
			doc:  fixtures[1],
			want: 1,
		},
		{
			name: "Тест 3",
			doc:  fixtures[2],
			want: 2,
		},
		{
			name: "Тест 4",
			doc:  fixtures[2],
			want: 2,
		},
		{
			name: "Тест 5",
			doc:  fixtures[1],
			want: 1,
		},
		{
			name: "Тест 6",
			doc:  fixtures[0],
			want: 0,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if gotId := tree.Insert(tt.doc); gotId != tt.want {
				t1.Errorf("Insert() = %v, want %v", gotId, tt.want)
			}
		})
	}
}

func TestTree_Find(t1 *testing.T) {
	tests := []struct {
		name    string
		id      int
		want    scanner.Document
		wantErr bool
	}{
		{
			name: "Тест1",
			id:   0,
			want: fixtures[0],
		},
		{
			name: "Тест2",
			id:   1,
			want: fixtures[1],
		},
		{
			name: "Тест3",
			id:   2,
			want: fixtures[2],
		},
		{
			name:    "Тест4",
			id:      3,
			wantErr: true,
		},
	}

	for _, d := range fixtures {
		tree.Insert(d)
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			got, err := tree.Find(tt.id)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !(got.URL == tt.want.URL && got.Title == tt.want.Title) {
				t1.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Serialize(t1 *testing.T) {
	for _, d := range fixtures {
		tree.Insert(d)
	}
	want := `[{"ID":0,"URL":"url1","Title":"заголовок из нескольких слов","Body":""},{"ID":1,"URL":"url2","Title":"иЗ нЕсКоЛЬких сЛОв","Body":""},{"ID":2,"URL":"url3","Title":"другой заголовок","Body":""}]`
	got, _ := tree.Json()
	if !reflect.DeepEqual(string(got), want) {
		t1.Errorf("Serialize() got = %s, want %s", string(got), want)
	}
}

func BenchmarkTree_InsertFind(b *testing.B) {
	s := Tree{}
	for i := 0; i < b.N; i++ {
		id := s.Insert(scanner.Document{
			URL:   "htt://url.com",
			Title: "Title",
		})
		_, _ = s.Find(id)
	}
}
