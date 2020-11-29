package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	var query string
	for {
		for query == "" {
			fmt.Printf("\nВведите строку для поиска: ")
			fmt.Scanln(&query)
		}
		_, err = conn.Write([]byte(query))
		if err != nil {
			log.Fatal(err)
		}

		data, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(data)
		query = ""
	}
}
