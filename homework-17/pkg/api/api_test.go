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
	user := User{Login: "admin", Pass: "P@ssw0rd", Admin: true}
	api := New(
		r,
		[]User{user},
		[]byte("trustno1"),
	)
	api.Endpoints()

	tests := []struct {
		name string
		user User
		want int
	}{
		{
			name: "Тест1",
			user: user,
			want: http.StatusOK,
		},
		{
			name: "Тест2",
			user: User{},
			want: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, _ := json.Marshal(tt.user)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/auth", bytes.NewBuffer(payload))
			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)
			if rr.Code != tt.want {
				t.Fatalf("код неверен: получили %d, а хотели %d", rr.Code, tt.want)
			}
			if rr.Code == http.StatusOK {
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
		})
	}
}
