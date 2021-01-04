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
	New(r)
	tests := []struct {
		name string
		user user
		want int
	}{
		{
			name: "Тест1",
			user: user{Login: "admin", Pass: "AdminP@ssw0rd", Admin: true},
			want: http.StatusOK,
		},
		{
			name: "Тест2",
			user: user{Login: "guest", Pass: "GuestP@ssw0rd", Admin: false},
			want: http.StatusOK,
		},
		{
			name: "Тест3",
			user: user{Login: "admin", Pass: "Wr0nG"},
			want: http.StatusUnauthorized,
		},
		{
			name: "Тест4",
			user: user{},
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
					return key, nil
				})
				if err != nil {
					t.Errorf("ошибка раскодирования токена: %v", err)
				}

				claims, ok := token.Claims.(jwt.MapClaims)
				if !(ok && token.Valid) {
					t.Error("токен невалидный")
				}
				if claims["admin"] != tt.user.Admin {
					t.Errorf("ошибка claims: получили %v, а хотели %v", claims["admin"], tt.user.Admin)
				}
			}
		})
	}
}
