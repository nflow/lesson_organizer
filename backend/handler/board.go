package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nflow/lesson_organizer/db"
)

func RetrieveBoards(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.BoardsMock)
}

func CreateBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func DeleteBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func RetrieveBoardGoals(w http.ResponseWriter, r *http.Request) {
	return
}

func CreateGoalInBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func UpdateGoalInBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func RemoveGoalFromBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func RetrieveBoardContents(w http.ResponseWriter, r *http.Request) {
	return
}

func CreateContentFromBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func UpdateContentFromBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func DeleteContentFromBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func RetrieveBoardPhases(w http.ResponseWriter, r *http.Request) {
	return
}

func AddPhaseToBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func UpdatePhaseInBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func RemovePhaseFromBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func RetrievePhaseMethods(w http.ResponseWriter, r *http.Request) {
	return
}

func AddMethodToPhase(w http.ResponseWriter, r *http.Request) {
	return
}

func UpdateMethodInPhase(w http.ResponseWriter, r *http.Request) {
	return
}

func DeleteMethodFromPhase(w http.ResponseWriter, r *http.Request) {
	return
}

func RetrieveMethodConents(w http.ResponseWriter, r *http.Request) {
	return
}

func AddContentToMethod(w http.ResponseWriter, r *http.Request) {
	return
}

func UpdateContentInMethod(w http.ResponseWriter, r *http.Request) {
	return
}

func RemoveContentFromMethod(w http.ResponseWriter, r *http.Request) {
	return
}
