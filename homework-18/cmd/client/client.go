package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	// авторизация и получение сообщений в фоне
	var client struct {
		Name string
		Pwd  string
	}
	for client.Name == "" {
		fmt.Printf("Введите имя: ")
		fmt.Scanln(&client.Name)
	}
	for client.Pwd == "" {
		fmt.Printf("Введите пароль: ")
		fmt.Scanln(&client.Pwd)
	}
	payload, err := json.Marshal(client)
	if err != nil {
		log.Fatalf("не удалось закодировать логин и пароль: %v", err)
	}

	ws, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/messages", nil)
	if err != nil {
		log.Fatalf("не удалось подключиться к серверу: %v", err)
	}
	err = ws.WriteMessage(websocket.TextMessage, payload)
	if err != nil {
		ws.Close()
		log.Fatalf("не удалось отправить сообщение: %v", err)
	}
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			ws.Close()
			log.Fatalf("не удалось прочитать сообщение: %v", err)
		}
		if string(p) == "OK" {
			break
		}
		ws.Close()
		log.Fatal("не удалось авторизоваться")
	}

	// получение сообщений в фоне
	go messages(ws)
	// интерактивная отправка сообщений в основном потоке
	send(client.Name)
}

func send(name string) {
	reader := bufio.NewReader(os.Stdin) // буфер для os.Stdin
	for {
		fmt.Print("-> ")
		msg, _ := reader.ReadString('\n')        // чтение строки (до символа перевода)
		msg = strings.Replace(msg, "\n", "", -1) // удаление перевода строки

		ws, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/send", nil)
		if err != nil {
			ws.Close()
			log.Fatalf("не удалось подключиться к серверу: %v", err)
		}

		err = ws.WriteMessage(websocket.TextMessage, []byte(name+": "+msg))
		if err != nil {
			ws.Close()
			log.Fatalf("не удалось отправить сообщение: %v", err)
		}
		ws.Close()
		time.Sleep(time.Second)
	}
}

func messages(ws *websocket.Conn) {
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			log.Fatalf("не удалось прочитать сообщение: %v", err)
		}
		fmt.Println(string(p))
		time.Sleep(time.Second)
	}
}
