package main

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nflow/lesson_organizer/db"
	"github.com/nflow/lesson_organizer/handler"
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	v := db.Method{
		uuid.New(),
		"Internetrecherche",
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
		"",
	}

	dbConnection.Create(v)

	router := mux.NewRouter()
	router.Use(commonMiddleware)

	router.HandleFunc("/v1/phases", handler.RetrievePhases).Methods("GET")
	router.HandleFunc("/v1/phases", handler.CreatePhase).Methods("POST")
	router.HandleFunc("/v1/phases/{id}", handler.DeletePhase).Methods("DELETE")

	router.HandleFunc("/v1/methods", handler.RetrieveMethods).Methods("GET")
	router.HandleFunc("/v1/methods", handler.CreateMethod).Methods("POST")
	router.HandleFunc("/v1/methods/{id}", handler.DeleteMethod).Methods("DELETE")

	router.HandleFunc("/v1/boards", handler.RetrieveBoards).Methods("GET")
	router.HandleFunc("/v1/boards", handler.CreateBoard).Methods("POST")
	router.HandleFunc("/v1/boards/{boardId}", handler.DeleteBoard).Methods("DELETE")

	router.HandleFunc("/v1/boards/{boardId}/goals", handler.RetrieveBoardGoals).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}/goals", handler.CreateGoalInBoard).Methods("POST")
	router.HandleFunc("/v1/boards/{boardId}/goals/{goalId}", handler.UpdateGoalInBoard).Methods("PUT")
	router.HandleFunc("/v1/boards/{boardId}/goals/{goalId}", handler.RemoveGoalFromBoard).Methods("DELETE")

	router.HandleFunc("/v1/boards/{boardId}/contents", handler.RetrieveBoardContents).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}/contents", handler.CreateContentFromBoard).Methods("POST")
	router.HandleFunc("/v1/boards/{boardId}/contents/{contentId}", handler.UpdateContentFromBoard).Methods("PUT")
	router.HandleFunc("/v1/boards/{boardId}/contents/{contentId}", handler.DeleteContentFromBoard).Methods("DELETE")

	router.HandleFunc("/v1/boards/{boardId}/phases", handler.RetrieveBoardPhases).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}/phases", handler.AddPhaseToBoard).Methods("POST")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}", handler.UpdatePhaseInBoard).Methods("PUT")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}", handler.RemovePhaseFromBoard).Methods("DELETE")

	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods", handler.RetrievePhaseMethods).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods", handler.AddMethodToPhase).Methods("POST")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods/{methodId}", handler.UpdateMethodInPhase).Methods("PUT")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods/{methodId}", handler.DeleteMethodFromPhase).Methods("DELETE")

	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods/{methodId}/contents", handler.RetrieveMethodConents).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods/{methodId}/contents", handler.AddContentToMethod).Methods("POST")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods/{methodId}/contents/{contentsId}", handler.UpdateContentInMethod).Methods("PUT")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods/{methodId}/contents/{contentsId}", handler.RemoveContentFromMethod).Methods("DELETE")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
