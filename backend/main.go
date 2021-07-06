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

	router.HandleFunc("/v1/phases", handler.RetrieveAllPhases).Methods("GET")
	router.HandleFunc("/v1/phases", handler.CreatePhase).Methods("POST")
	router.HandleFunc("/v1/phases/{id}", handler.DeletePhase).Methods("DELETE")

	router.HandleFunc("/v1/methods", handler.RetrieveAllMethods).Methods("GET")
	router.HandleFunc("/v1/methods", handler.CreateMethod).Methods("POST")
	router.HandleFunc("/v1/methods/{id}", handler.DeleteMethod).Methods("DELETE")

	router.HandleFunc("/v1/boards", handler.RetrieveAllBoards).Methods("GET")
	router.HandleFunc("/v1/boards", handler.CreateBoard).Methods("POST")
	router.HandleFunc("/v1/boards/{boardId}", handler.DeleteBoard).Methods("DELETE")

	router.HandleFunc("/v1/boards/{boardId}/goals", handler.RetrieveBoardGoals).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}/goals/{goalId}", handler.RetrieveBoardGoals).Methods("GET")

	router.HandleFunc("/v1/boards/{boardId}/phases", handler.RetrieveBoardPhases).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}", handler.RetrieveBoardPhases).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods", handler.RetrieveBoardPhases).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods/{methodId}", handler.RetrieveBoardPhases).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods/{methodId}/contents", handler.RetrieveBoardPhases).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods/{methodId}/contents/{contentsId}", handler.RetrieveBoardPhases).Methods("GET")

	router.HandleFunc("/v1/boards/{boardId}/contents", handler.RetrieveBoardContents).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}/contents/{contentId}", handler.RetrieveBoardContents).Methods("GET")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
