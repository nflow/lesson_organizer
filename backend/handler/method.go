package handler

import (
	"errors"
	"net/http"

	"github.com/nflow/lesson_organizer/model"
	"gorm.io/gorm"
)

func (h *Handler) RetrieveMethods(w http.ResponseWriter, r *http.Request) {
	var method model.Method

	if err := h.DB.Take(&method).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		RespondWithEmptyArray(w)

		return
	} else if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	RespondWithSuccess(w, method)
}

func (h *Handler) CreateMethod(w http.ResponseWriter, r *http.Request) {
	var method model.Method

	if !HandleBodyDecode(w, r, &method) {
		return
	}

	RespondWithCode(w, http.StatusCreated, method)
}

func (h *Handler) DeleteMethod(w http.ResponseWriter, r *http.Request) {
	return
}
