package api

import (
	"bytes"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI_authHandler(t *testing.T) {
	r := mux.NewRouter()
	api := New(r)

	user := user{
		Login: "admin",
		Pass:  "P@ssw0rd",
		Admin: true,
	}
	payload, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {

	}
	resp, _ := ioutil.ReadAll(rr.Body)
	token, err := jwt.Parse(string(resp), func(token *jwt.Token) (interface{}, error) {
		return api.key, nil
	})
	if err != nil {
		t.Errorf("ошибка раскодирования токена: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		t.Error("токен невалидный")
	}
	if claims["admin"] != user.Admin {
		t.Errorf("ошибка claims: получили %v, а хотели %v", claims["admin"], user.Admin)
	}
}
