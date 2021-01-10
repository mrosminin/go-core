// RPC клиент для поисковика
package main

import (
	"fmt"
	"go-core-own/homework-19/pkg/scanner"
	"log"
	"net/rpc"
)

const addr = ":20001"

func main() {
	var query string
	for {
		for query == "" {
			fmt.Printf("\nВведите строку для поиска: ")
			fmt.Scanln(&query)
		}

		// создание клиента RPC
		client, err := rpc.Dial("tcp", addr)
		if err != nil {
			log.Fatal("dialing:", err)
		}
		var docs []scanner.Document
		client.Call("RPC.Search", query, &docs)
		if err != nil {
			log.Fatal(err)
		}
		client.Close()
		for i, d := range docs {
			fmt.Printf("%d %v\n", i+1, d)
		}
		query = ""
	}
}
