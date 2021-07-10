package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func HandleBodyDecode(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		RespondWithError(w, http.StatusBadRequest, err)
		return false
	}
	defer r.Body.Close()

	return true
}

func RespondWithError(w http.ResponseWriter, code int, err error) {
	log.Println(err)
	RespondWithCode(w, code, map[string]string{"error": err.Error()})
}

func RespondWithSuccess(w http.ResponseWriter, v interface{}) {
	RespondWithCode(w, http.StatusOK, v)
}

func RespondWithCode(w http.ResponseWriter, code int, v interface{}) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}

func RespondWithEmptyArray(w http.ResponseWriter) {
	json.NewEncoder(w).Encode([]string{})
}
