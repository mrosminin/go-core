package api

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

type User struct {
	Login string
	Pass  string
	Admin bool
}

type API struct {
	r     *mux.Router
	users []User
	key   []byte
}

func New(r *mux.Router, users []User, key []byte) *API {
	return &API{
		r:     r,
		users: users,
		key:   key,
	}
}

func (api *API) Endpoints() {
	api.r.Use(logMiddleware)
	api.r.HandleFunc("/api/v1/auth", api.authHandler).Methods(http.MethodPost, http.MethodOptions)
}

// middleware для логирования запросов к API в формете Apache Common Log Format (CLF)
// http://httpd.apache.org/docs/2.2/logs.html#common
func logMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

func (api *API) authHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == http.MethodOptions {
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, u := range api.users {
		if user.Login == u.Login && user.Pass == u.Pass {
			user = u
			break
		}
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":   time.Now().Unix(),
		"admin": user.Admin,
	})

	tokenString, err := token.SignedString(api.key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(tokenString))
}
