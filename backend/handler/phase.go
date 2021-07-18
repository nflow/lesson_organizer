package handler

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nflow/lesson_organizer/model"
	"gorm.io/gorm"
)

func (h *Handler) RetrievePhases(w http.ResponseWriter, r *http.Request) {
	Phase := &[]model.Phase{}

	if err := h.DB.Find(&Phase).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		RespondWithEmptyArray(w)

		return
	} else if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err)

		return
	}

	RespondWithSuccess(w, Phase)
}

func (h *Handler) CreatePhase(w http.ResponseWriter, r *http.Request) {
	Phase := &model.Phase{}

	if !HandleBodyDecode(w, r, Phase) {
		return
	}

	Phase.ID = uuid.New()

	if err := h.DB.Create(Phase).Error; err != nil {
		RespondWithError(w, http.StatusInternalServerError, err)
	}

	RespondWithCode(w, http.StatusCreated, Phase)
}

func (h *Handler) ModifyPhase(w http.ResponseWriter, r *http.Request) {
	Phase := &model.Phase{}

	if !HandleBodyDecode(w, r, Phase) {
		return
	}

	if err := h.DB.Save(&Phase).Error; err != nil {
		RespondWithError(w, http.StatusNotFound, err)
		return
	}

	RespondWithSuccess(w, Phase)
}

func (h *Handler) DeletePhase(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Phase := model.Phase{}

	if err := h.DB.First(&Phase, "id = ?", vars["id"]).Delete(&Phase).Error; err != nil {
		RespondWithError(w, http.StatusNotFound, err)
		return
	}

	RespondWithSuccess(w, Phase)
}
