package main

import (
	"flag"
	"fmt"
	"go-core-own/pkg/fibo"
	"log"
	"strconv"
)

func main() {
	var nFlag = flag.Int("n", 0, "Номер числа Фибоначчи")
	flag.Parse()
	n := *nFlag

	if n == 0 {
		var input string
		fmt.Printf("Укажите номер числа Фибоначчи: ")
		fmt.Scanln(&input)
		var err error
		n, err = strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
	}

	if n > 20 {
		log.Fatal("Номер должен быть менее 20")
	}

	var result int = fibo.Fibo(n)
	fmt.Printf("Число Фибоначчи номер %v: %v\n", n, result)
}
