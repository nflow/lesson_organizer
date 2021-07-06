package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nflow/lesson_organizer/model"
)

func RetrieveMethods(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.MethodsMock)
}

func CreateMethod(w http.ResponseWriter, r *http.Request) {
	return
}

func DeleteMethod(w http.ResponseWriter, r *http.Request) {
	return
}
