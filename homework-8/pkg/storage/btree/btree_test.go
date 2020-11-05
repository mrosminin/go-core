package btree

import (
	"go-core-own/homework-8/pkg/scanner"
	"testing"
)

func TestTree(t *testing.T) {
	tree := Tree{}
	fixtures := []scanner.Document{
		{URL: "url1", Title: "заголовок из нескольких слов"},
		{URL: "url2", Title: "иЗ нЕсКоЛЬких сЛОв"},
		{URL: "url3", Title: "_нескольких_слов_"},
	}

	for i := 0; i < len(fixtures); i++ {
		fixtures[i].ID = tree.Insert(fixtures[i])
	}

	for _, d := range fixtures {
		got, err := tree.Find(d.ID)
		if err != nil {
			t.Errorf("Ошибка %v:", err)
		}
		if got.ID != d.ID {
			t.Errorf("BTree got %d, want %d", got.ID, d.ID)
		}
	}
}
