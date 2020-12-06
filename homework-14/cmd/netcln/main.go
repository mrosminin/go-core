package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-core-own/homework-14/pkg/scanner"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Введите запрос: ")
		query, _ := reader.ReadString('\n')
		if _, err = fmt.Fprintf(conn, "%s\n", query); err != nil {
			log.Fatal(err)
		}
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}

		var docs []scanner.Document
		err = json.Unmarshal([]byte(response), &docs)
		if err != nil {
			fmt.Println(response)
			continue
		}
		for i, d := range docs {
			fmt.Printf("%d %v\n", i+1, d)
		}
	}
}
