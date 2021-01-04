package main

import (
	"github.com/gorilla/mux"
	"go-core-own/homework-17/pkg/api"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	api.New(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
