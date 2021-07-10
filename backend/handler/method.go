package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nflow/lesson_organizer/model"
	"gorm.io/gorm"
)

func (h *Handler) RetrieveMethods(w http.ResponseWriter, r *http.Request) {
	method := &[]model.Method{}

	if err := h.DB.Find(&method).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		RespondWithEmptyArray(w)

		return
	} else if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err)

		return
	}

	RespondWithSuccess(w, method)
}

func (h *Handler) CreateMethod(w http.ResponseWriter, r *http.Request) {
	method := &model.Method{}

	if !HandleBodyDecode(w, r, method) {
		return
	}

	method.ID = uuid.New()

	if err := h.DB.Create(method).Error; err != nil {
		RespondWithError(w, http.StatusInternalServerError, err)
	}

	RespondWithCode(w, http.StatusCreated, method)
}

func (h *Handler) DeleteMethod(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	log.Println(vars)

	method := model.Method{}
	if err := h.DB.First(&method, "id = ?", vars["id"]).Error; err != nil {
		RespondWithError(w, http.StatusNotFound, err)
		return
	}

	if err := h.DB.Delete(&method).Error; err != nil {
		RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	RespondWithSuccess(w, method)
}
