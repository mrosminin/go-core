package api

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"go-core-own/homework-16/pkg/engine"
	"go-core-own/homework-16/pkg/index"
	"go-core-own/homework-16/pkg/scanner"
	"go-core-own/homework-16/pkg/storage"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

var api *API

func TestMain(m *testing.M) {
	api = &API{
		r: mux.NewRouter(),
		e: engine.New(index.New(), &storage.Service{}),
	}
	api.endpoints()
	os.Exit(m.Run())
}

func TestAPI_FindRequestHandler(t *testing.T) {
	want := scanner.Document{ID: 0, URL: "url1", Title: "ЗаГолоВОК иЗ нЕсКоЛЬких сЛОв"}
	api.e.Store([]scanner.Document{want})

	req := httptest.NewRequest(http.MethodPost, "/api/public/v1/find", bytes.NewBuffer([]byte("заголовок")))
	rr := httptest.NewRecorder()
	api.r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	var got []scanner.Document
	resp, _ := ioutil.ReadAll(rr.Body)
	_ = json.Unmarshal(resp, &got)

	if !reflect.DeepEqual(want, got[0]) {
		t.Errorf("FindRequestHandler() = %v, want %v", got[0], want)
	}
}

func TestAPI_NewDocRequestHandler(t *testing.T) {
	doc := scanner.Document{ID: 0, URL: "url1", Title: "ЗаГолоВОК иЗ нЕсКоЛЬких сЛОв"}
	payload, _ := json.Marshal(doc)

	wantLen := len(api.e.Storage.Docs) + 1

	req := httptest.NewRequest(http.MethodPost, "/api/public/v1/newDoc", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	if gotLen := len(api.e.Storage.Docs); wantLen != gotLen {
		t.Fatalf("NewDocRequestHandler() gotLen %d, wantLen %d", gotLen, wantLen)
	}
	if got := api.e.Storage.Docs[len(api.e.Storage.Docs)-1]; !reflect.DeepEqual(doc, got) {
		t.Errorf("NewDocRequestHandler() got %v, want %v", got, doc)
	}
}

func TestAPI_UpdateDocRequestHandler(t *testing.T) {
	doc := scanner.Document{ID: 0, URL: "url1", Title: "ЗаГолоВОК иЗ нЕсКоЛЬких сЛОв"}
	api.e.Store([]scanner.Document{doc})

	doc.URL = "url2"
	payload, _ := json.Marshal(doc)

	req := httptest.NewRequest(http.MethodPost, "/api/public/v1/updateDoc", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	if got := api.e.Storage.Docs[len(api.e.Storage.Docs)-1]; doc.URL != got.URL {
		t.Errorf("UpdateDocRequestHandler() got %v, want %v", got, doc)
	}
}

func TestAPI_DeleteDocRequestHandler(t *testing.T) {
	doc := scanner.Document{ID: 0, URL: "url1", Title: "ЗаГолоВОК иЗ нЕсКоЛЬких сЛОв"}
	api.e.Store([]scanner.Document{doc})

	wantLen := len(api.e.Storage.Docs) - 1

	payload, _ := json.Marshal(doc)

	req := httptest.NewRequest(http.MethodPost, "/api/public/v1/deleteDoc", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	if gotLen := len(api.e.Storage.Docs); wantLen != gotLen {
		t.Fatalf("NewDocRequestHandler() gotLen %d, wantLen %d", gotLen, wantLen)
	}
}
