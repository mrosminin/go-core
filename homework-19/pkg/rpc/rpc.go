package rpc

import (
	"go-core-own/homework-19/pkg/engine"
	"go-core-own/homework-19/pkg/scanner"
	"log"
	"net"
	"net/rpc"
)

type RPC struct {
	e *engine.Service
}

type Service struct {
	rpc *RPC
}

func New(e *engine.Service) *Service {
	return &Service{&RPC{e}}
}

func (s *Service) Init(addr string) error {
	err := rpc.Register(s.rpc)
	if err != nil {
		return err
	}
	// регистрация сетевой службы RPC-сервера
	listener, err := net.Listen("tcp4", addr)
	if err != nil {
		return err
	}
	// цикл обработки клиентских подключений
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			go rpc.ServeConn(conn)
		}
	}()
	return nil
}

func (rpc *RPC) Search(query string, result *[]scanner.Document) error {
	docs := rpc.e.Find(query)
	*result = docs
	return nil
}
