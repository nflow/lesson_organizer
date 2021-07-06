package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nflow/lesson_organizer/model"
)

func MethodHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func RetrieveAllMethods(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.MethodsMock)
}
