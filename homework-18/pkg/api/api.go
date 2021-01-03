package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type Client struct {
	Name  string
	Queue chan string
}

type API struct {
	r        *mux.Router
	mux      *sync.Mutex
	upgrader websocket.Upgrader
	Queue    chan string
	Clients  []Client
}

func New(r *mux.Router) *API {
	return &API{
		r:       r,
		mux:     &sync.Mutex{},
		Queue:   make(chan string),
		Clients: make([]Client, 0),
	}
}

func (api *API) Endpoints() {
	api.r.HandleFunc("/send", api.sendHandler)
	api.r.HandleFunc("/messages", api.messagesHandler)
}

// приём сообщений от клиентов
func (api *API) sendHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := api.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	mt, message, err := conn.ReadMessage()
	if err != nil {
		conn.WriteMessage(mt, []byte(err.Error()))
		return
	}
	fmt.Println("получено сообщение:", string(message))

	// все входящие сообщения пишутся в очередь,
	// дальше они обрабатываются в потоке publishMessages
	api.Queue <- string(message)
}

// получение сообщений от всех клиентов
func (api *API) messagesHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := api.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	_, message, err := conn.ReadMessage()
	if err != nil {
		conn.WriteMessage(websocket.CloseMessage, []byte(err.Error()))
		return
	}

	var req struct {
		Name string
		Pwd  string
	}
	err = json.Unmarshal(message, &req)
	if err != nil {
		conn.WriteMessage(websocket.CloseMessage, []byte(err.Error()))
		return
	}
	if req.Pwd != "password" {
		conn.WriteMessage(websocket.TextMessage, []byte("Unauthorized"))
		return
	}

	// при подключении для клиента создаётся канал и добавляется в массив
	api.mux.Lock()
	client := Client{
		Name:  req.Name,
		Queue: make(chan string),
	}
	api.Clients = append(api.Clients, client)
	api.mux.Unlock()

	conn.WriteMessage(websocket.TextMessage, []byte("OK"))

	// при отключении канал удаляется из массива, чтобы избежать паники.
	defer func() {
		api.mux.Lock()
		for i := range api.Clients {
			if api.Clients[i] == client {
				api.Clients = append(api.Clients[:i], api.Clients[i+1:]...)
				break
			}
		}
		api.mux.Unlock()
	}()

	// чтение сообщений из канала данного клиента
	for msg := range client.Queue {
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			return
		}
	}
}
