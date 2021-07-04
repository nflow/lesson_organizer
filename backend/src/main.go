package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type MethodLabel string

const (
	METHOD_LABEL_SINGLE = "METHOD_LABEL_SINGLE"
	METHOD_LABEL_PAIR   = "METHOD_LABEL_PAIR"
	METHOD_LABEL_GROUP  = "METHOD_LABEL_GROUP"
	METHOD_LABEL_PLENUM = "METHOD_LABEL_PLENUM"
)

type Board struct {
	Id   uuid.UUID `json:id`
	Name string
}

type Goal struct {
	Id   uuid.UUID `json:id`
	Text string    `json:text`
}

type Phase struct {
	Id    uuid.UUID `json:id`
	Title string    `json:title`
}

type Method struct {
	Id          uuid.UUID     `json:id`
	Title       string        `json:title`
	Description string        `json:description`
	Labels      []MethodLabel `json:labels`
}

type Content struct {
	Id    uuid.UUID `json:id`
	Title string    `json:title`
	Text  string    `json:text`
}

var methods = []Method{
	{
		uuid.New(),
		"Mind-Map",
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
		nil,
	},
	{
		uuid.New(),
		"Blitzlicht",
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
		nil,
	},
	{
		uuid.New(),
		"Soziometrische Abfrage",
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
		nil,
	},
	{
		uuid.New(),
		"Internetrecherche",
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
		nil,
	},
}

var boards = []Board{
	{
		uuid.New(),
		"Test Board",
	},
}

func retrieveAllMethods(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(methods)
}

func retrieveAllBoards(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(boards)
}

func retrieveBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	for _, v := range boards {
		if vars["id"] == v.Id.String() {
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.Use(commonMiddleware)

	router.HandleFunc("/v1/methods", retrieveAllMethods).Methods("GET")
	router.HandleFunc("/v1/boards", retrieveAllBoards).Methods("GET")

	router.HandleFunc("/v1/board/{id}", retrieveBoard).Methods("GET")
	router.HandleFunc("/v1/goal/{id}", retrieveBoard).Methods("GET")
	router.HandleFunc("/v1/phase/{id}", retrieveBoard).Methods("GET")
	router.HandleFunc("/v1/method/{id}", retrieveBoard).Methods("GET")
	router.HandleFunc("/v1/content/{id}", retrieveBoard).Methods("GET")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
