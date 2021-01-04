package rpc

import (
	"go-core-own/homework-19/pkg/engine"
	"go-core-own/homework-19/pkg/index"
	"go-core-own/homework-19/pkg/scanner"
	"go-core-own/homework-19/pkg/storage"
	"reflect"
	"testing"
)

func TestRPC_Search(t *testing.T) {
	rpc := &RPC{}
	rpc.e = engine.New(index.New(), &storage.Service{})

	want := scanner.Document{ID: 0, URL: "url1", Title: "ЗаГолоВОК иЗ нЕсКоЛЬких сЛОв"}
	rpc.e.Store([]scanner.Document{want})

	var got []scanner.Document

	_ = rpc.Search("заголовок", &got)
	if !reflect.DeepEqual(want, got[0]) {
		t.Errorf("rpc.Search() = %v, want %v", got[0], want)
	}
}
