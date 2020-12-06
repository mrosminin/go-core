package webapp

import (
	"fmt"
	"go-core-own/homework-15/pkg/index"
	"go-core-own/homework-15/pkg/storage"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	network = "tcp4"
	addr    = ":8080"
)

type Service struct {
	srv *http.Server

	index   *index.Service
	storage *storage.Service
}

func New(index *index.Service, storage *storage.Service) *Service {

	// параметры веб-сервера
	srv := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		Handler:      nil,
		Addr:         addr,
	}

	return &Service{
		srv:     srv,
		index:   index,
		storage: storage,
	}
}

func (s *Service) Init() error {
	// старт сетевой службы веб-сервера
	listener, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}

	// регистрация обработчика для URL `/` в маршрутизаторе по умолчанию
	http.HandleFunc("/index", s.indexHandler)
	http.HandleFunc("/docs", s.docsHandler)
	http.HandleFunc("/", s.mainHandler)

	// старт самого веб-сервера
	err = s.srv.Serve(listener)
	return err
}

func (s *Service) mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body>")
	fmt.Fprintf(w, "<a href=\"http://%s/index\">Index</a><br>", r.Host)
	fmt.Fprintf(w, "<a href=\"http://%s/docs\">Docs</a>", r.Host)
	fmt.Fprint(w, "</body></html>")
}

func (s *Service) indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body><div>")
	fmt.Fprintf(w, "%v", s.index.Index)
	fmt.Fprint(w, "</div></body></html>")
}

func (s *Service) docsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body><div>")
	defer fmt.Fprint(w, "</div></body></html>")
	str, err := s.storage.Json()
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	fmt.Fprintf(w, "%s", str)
}
