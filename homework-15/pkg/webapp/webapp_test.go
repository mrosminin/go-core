package webapp

import (
	"encoding/json"
	"go-core-own/homework-15/pkg/index"
	"go-core-own/homework-15/pkg/storage"
	"go-core-own/homework-15/pkg/storage/btree"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestService_mainHandler(t *testing.T) {
	s := &Service{}
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	rr := httptest.NewRecorder()
	s.mainHandler(rr, req)

	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	resp := rr.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var data string
	json.Unmarshal(body, &data)
	t.Logf("Ответ сервера:\n%v\n", data)

}

func TestService_indexHandler(t *testing.T) {
	s := &Service{
		index: index.New(),
	}
	req := httptest.NewRequest(http.MethodGet, "/index", nil)

	rr := httptest.NewRecorder()
	s.indexHandler(rr, req)

	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	resp := rr.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var data string
	json.Unmarshal(body, &data)
	t.Logf("Ответ сервера:\n%v\n", data)

}

func TestService_docsHandler(t *testing.T) {
	s := &Service{
		storage: storage.New(nil, btree.New()),
	}
	req := httptest.NewRequest(http.MethodGet, "/docs", nil)

	rr := httptest.NewRecorder()
	s.docsHandler(rr, req)

	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	resp := rr.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var data string
	json.Unmarshal(body, &data)
	t.Logf("Ответ сервера:\n%v\n", data)
}
