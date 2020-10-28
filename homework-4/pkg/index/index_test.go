package index

import (
	"testing"
)

func Test_Find(t *testing.T) {
	i := New()
	fixtures := make(map[string]string)
	fixtures["url"] = "заголовок из нескольких слов"
	i.Fill(fixtures)
	want := 1
	got := i.Find("заголовок")
	if l := len(got); l != want {
		t.Errorf("Find(): получили %d, должны были %d\n", l, want)
	}
}
