// Групповой чат по WebSocket - сервер
package main

import (
	"github.com/gorilla/mux"
	"go-core-own/homework-18/pkg/api"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	api := api.New(r)
	api.Endpoints()

	// отправка сообщений всем подключенным клиентам
	// Поступающие сообщения записываются в канал Queue,
	// откуда они перенаправляются в каналы-клиенты.
	// Шаблон Fan-Out.
	go func() {
		for msg := range api.Queue {
			for _, c := range api.Clients {
				c.Queue <- msg
			}
		}
	}()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
