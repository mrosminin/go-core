package rpc

import (
	"go-core-own/homework-19/pkg/engine"
	"go-core-own/homework-19/pkg/index"
	"go-core-own/homework-19/pkg/scanner"
	"go-core-own/homework-19/pkg/storage"
	"log"
	"net"
	"net/rpc"
	"reflect"
	"testing"
)

const addr = ":20001"

func TestRPC_Search(t *testing.T) {
	s := new(RPC)
	s.e = engine.New(index.New(), &storage.Service{})
	want := scanner.Document{ID: 0, URL: "url1", Title: "ЗаГолоВОК иЗ нЕсКоЛЬких сЛОв"}
	s.e.Store([]scanner.Document{want})

	err := rpc.Register(s)
	if err != nil {
		t.Fatal("registering rpc service", err)
	}
	err = rpc.RegisterName("RPC.Search", s)
	if err != nil {
		t.Fatal("registering rpc method", err)
	}
	l, err := net.Listen("tcp4", addr)
	if err != nil {
		t.Fatal("listening", err)
	}

	go func() {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go rpc.ServeConn(conn)
	}()

	client, err := rpc.Dial("tcp4", addr)
	if err != nil {
		t.Fatal("dialing", err)
	}
	defer client.Close()

	var got []scanner.Document
	err = client.Call("RPC.Search", "заголовок", &got)

	if err != nil {
		t.Errorf("Add: expected no error but got string %q", err.Error())
	}

	if !(len(got) > 0 && reflect.DeepEqual(want, got[0])) {
		t.Errorf("rpc.Search() = %v, want %v", got, want)
	}
}
