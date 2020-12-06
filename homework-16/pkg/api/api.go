package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-core-own/homework-16/pkg/engine"
	"go-core-own/homework-16/pkg/scanner"
	"net/http"
)

type API struct {
	r      *mux.Router
	engine *engine.Service
}

func New(e *engine.Service) *API {
	api := API{
		r:      mux.NewRouter(),
		engine: e,
	}
	return &api
}

func (api *API) Init(addr string) error {
	api.endpoints()
	err := http.ListenAndServe(addr, api.r)
	if err != nil {
		return err
	}
	return nil
}

func (api *API) endpoints() {
	api.r.HandleFunc("/api/public/v1/find", api.FindRequestHandler).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/api/public/v1/newDocs", api.NewDocsRequestHandler).Methods(http.MethodPost, http.MethodOptions)
}

func (api *API) FindRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, SessionID")

	if r.Method == http.MethodOptions {
		return
	}

	var query string
	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	docs := api.engine.Find(query)
	jsonData, err := json.Marshal(docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (api *API) NewDocsRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, SessionID")

	if r.Method == http.MethodOptions {
		return
	}

	var docs []scanner.Document
	err := json.NewDecoder(r.Body).Decode(&docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.engine.Store(docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
