package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nflow/lesson_organizer/model"
)

func BoardHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func RetrieveAllBoards(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.BoardsMock)
}

func RetrieveBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	for _, v := range model.BoardsMock {
		if vars["id"] == v.Id.String() {
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func RetrieveBoardGoals(w http.ResponseWriter, r *http.Request) {
	return
}

func RetrieveBoardPhases(w http.ResponseWriter, r *http.Request) {
	return
}

func RetrieveBoardContents(w http.ResponseWriter, r *http.Request) {
	return
}
