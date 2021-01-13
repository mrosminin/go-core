package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-core-own/homework-10/pkg/engine"
	"go-core-own/homework-10/pkg/scanner"
	"io/ioutil"
	"net/http"
	"net/http/pprof"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	searchRequestsNumber = promauto.NewCounter(prometheus.CounterOpts{
		Name: "search_requests_total",
		Help: "Количество поисковых запросов.",
	})
	searchRequestsTime = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "search_request_time",
		Help:    "Время выполнения поискового запроса, мс.",
		Buckets: prometheus.LinearBuckets(10, 10, 20), // 20 корзин, начиная с 0, по 10 элементов
	})
	searchRequestsAverageLength = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "search_requests_average_length",
		Help: "Средняя длина поискового запроса, байт.",
	})

	mutex                     = sync.Mutex{}
	searchRequestsCounter     = 0
	searchRequestsTotalLength = 0
)

type API struct {
	r *mux.Router
	e *engine.Service
}

func New(e *engine.Service) *API {
	api := API{
		r: mux.NewRouter(),
		e: e,
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
	// метрики Prometheus
	api.r.Handle("/metrics", promhttp.Handler())

	// профилирование приложения с помощью pprof
	api.r.HandleFunc("/debug/pprof/", pprof.Index)
	api.r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	api.r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	api.r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	api.r.HandleFunc("/debug/pprof/trace", pprof.Trace)
	api.r.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))
	api.r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	api.r.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	api.r.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	api.r.Handle("/debug/pprof/block", pprof.Handler("block"))

	api.r.HandleFunc("/api/public/v1/find", api.FindRequestHandler).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/api/public/v1/newDoc", api.NewDocRequestHandler).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/api/public/v1/updateDoc", api.UpdateDocRequestHandler).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/api/public/v1/deleteDoc", api.DeleteDocRequestHandler).Methods(http.MethodPost, http.MethodOptions)
}

func (api *API) FindRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, SessionID")

	if r.Method == http.MethodOptions {
		return
	}

	query, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t := time.Now()

	docs := api.e.Find(string(query))
	jsonData, err := json.Marshal(docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dur := time.Since(t).Milliseconds()
	searchRequestsTime.Observe(float64(dur))
	searchRequestsNumber.Inc()

	mutex.Lock()
	searchRequestsCounter++
	searchRequestsTotalLength += len(query)
	searchRequestsAverageLength.Set(float64(searchRequestsTotalLength / searchRequestsCounter))
	mutex.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (api *API) NewDocRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	if r.Method == http.MethodOptions {
		return
	}

	var doc scanner.Document
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	api.e.Store([]scanner.Document{doc})
	w.WriteHeader(http.StatusOK)
}

func (api *API) UpdateDocRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	if r.Method == http.MethodOptions {
		return
	}

	var doc scanner.Document
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.e.Storage.Update(doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) DeleteDocRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	if r.Method == http.MethodOptions {
		return
	}

	var doc scanner.Document
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.e.Storage.Delete(doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
