package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nflow/lesson_organizer/handler"
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.Use(commonMiddleware)

	router.HandleFunc("/v1/methods", handler.RetrieveAllMethods).Methods("GET")
	router.HandleFunc("/v1/boards", handler.RetrieveAllBoards).Methods("GET")

	router.HandleFunc("/v1/board/{id}", handler.BoardHandler)
	router.HandleFunc("/v1/board/{id}/goals", handler.RetrieveBoardGoals).Methods("GET")
	router.HandleFunc("/v1/board/{id}/phases", handler.RetrieveBoardPhases).Methods("GET")
	router.HandleFunc("/v1/board/{id}/contents", handler.RetrieveBoardContents).Methods("GET")

	router.HandleFunc("/v1/goal/{id}", handler.GoalHandler)

	router.HandleFunc("/v1/phase/{id}", handler.PhaseHandler)
	router.HandleFunc("/v1/phase/{id}/methods", handler.RetrieveMethodsFromPhase)

	router.HandleFunc("/v1/method/{id}/contents", handler.RetrieveAllContentsFromPhase)

	router.HandleFunc("/v1/content/{id}", handler.ContentHandler)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
