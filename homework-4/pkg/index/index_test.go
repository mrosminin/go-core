package index

import (
	"testing"
)

func Test_Find(t *testing.T) {
	i := New()
	fixtures := map[string]string{
		"url1": "заголовок из нескольких слов",
		"url2": "из нескольких слов",
	}
	i.Fill(fixtures)

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
			if got := i.Find(tt.s); len(got) != tt.want {
				t.Errorf("Find('%s') = %v, want %v", tt.s, len(got), tt.want)
			}
		})
	}
}
