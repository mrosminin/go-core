package main

import (
	"github.com/gorilla/mux"
	"go-core-own/homework-17/pkg/api"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	api := api.New(r)
	api.Endpoints()
	http.ListenAndServe(":8080", r)
}
