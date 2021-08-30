package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
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

	h := &handler.Handler{
		DB: dbConnection,
	}

	router := mux.NewRouter()
	router.Use(commonMiddleware)
	router.Use(handlers.CORS())

	router.HandleFunc("/v1/phases", h.RetrievePhases).Methods("GET")
	router.HandleFunc("/v1/phases", h.CreatePhase).Methods("POST")
	router.HandleFunc("/v1/phases/{id}", h.ModifyPhase).Methods("PUT")
	router.HandleFunc("/v1/phases/{id}", h.DeletePhase).Methods("DELETE")

	router.HandleFunc("/v1/methods", h.RetrieveMethods).Methods("GET")
	router.HandleFunc("/v1/methods", h.CreateMethod).Methods("POST")
	router.HandleFunc("/v1/methods/{id}", h.ModifyMethod).Methods("PUT")
	router.HandleFunc("/v1/methods/{id}", h.DeleteMethod).Methods("DELETE")

	router.HandleFunc("/v1/boards", h.RetrieveBoards).Methods("GET")
	router.HandleFunc("/v1/boards", h.CreateBoard).Methods("POST")
	router.HandleFunc("/v1/boards/{boardId}", h.RetrieveBoard).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}", h.DeleteBoard).Methods("DELETE")

	router.HandleFunc("/v1/boards/{boardId}/goals", h.RetrieveBoardGoals).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}/goals", h.CreateGoalInBoard).Methods("POST")
	router.HandleFunc("/v1/boards/{boardId}/goals/{goalId}", h.UpdateGoalInBoard).Methods("PUT")
	router.HandleFunc("/v1/boards/{boardId}/goals/{goalId}", h.RemoveGoalFromBoard).Methods("DELETE")

	router.HandleFunc("/v1/boards/{boardId}/contents", h.RetrieveBoardContents).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}/contents", h.AddContentToBoard).Methods("POST")
	router.HandleFunc("/v1/boards/{boardId}/contents/{contentId}", h.UpdateContentFromBoard).Methods("PUT")
	router.HandleFunc("/v1/boards/{boardId}/contents/{contentId}", h.DeleteContentFromBoard).Methods("DELETE")

	router.HandleFunc("/v1/boards/{boardId}/phases", h.RetrieveBoardPhases).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}/phases", h.AddPhaseToBoard).Methods("POST")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}", h.UpdatePhaseInBoard).Methods("PUT")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}", h.RemovePhaseFromBoard).Methods("DELETE")

	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods", handler.RetrievePhaseMethods).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods", h.AddMethodToPhase).Methods("POST")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods/{methodId}", handler.UpdateMethodInPhase).Methods("PUT")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods/{methodId}", handler.DeleteMethodFromPhase).Methods("DELETE")

	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods/{methodId}/contents", handler.RetrieveMethodConents).Methods("GET")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods/{methodId}/contents", h.AddContentToMethod).Methods("POST")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods/{methodId}/contents/{contentsId}", handler.UpdateContentInMethod).Methods("PUT")
	router.HandleFunc("/v1/boards/{boardId}/phases/{phaseId}/methods/{methodId}/contents/{contentsId}", handler.RemoveContentFromMethod).Methods("DELETE")

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	srv := &http.Server{
		Handler:      handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(loggedRouter),
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
