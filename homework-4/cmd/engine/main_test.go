package main

import (
	"go-core-own/homework-4/pkg/crawler/membot"
	"go-core-own/homework-4/pkg/index"
	"testing"
)

func Test_Engine(t *testing.T) {
	e := Engine{
		crawler: membot.New(),
		index:   index.New(),
	}
	e.ScanPage("", depth)

	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Тест №1",
			s:    "яНДЕКС",
			want: 1,
		},
		{
			name: "Тест №2",
			s:    "google",
			want: 1,
		},
		{
			name: "Тест №3",
			s:    "yahoo",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := e.Find(tt.s); len(got) != tt.want {
				t.Errorf("Engine('%s') got %v, want %v", tt.s, len(got), tt.want)
			}
		})
	}
}
