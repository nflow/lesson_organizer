package handler

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nflow/lesson_organizer/model"
	"gorm.io/gorm"
)

func (h *Handler) RetrieveBoards(w http.ResponseWriter, r *http.Request) {
	var board []model.Board

	if err := h.DB.Find(&board).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		RespondWithEmptyArray(w)

		return
	} else if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err)

		return
	}

	RespondWithSuccess(w, board)
}

func (h *Handler) CreateBoard(w http.ResponseWriter, r *http.Request) {
	board := &model.Board{}

	if !HandleBodyDecode(w, r, board) {
		return
	}
	board.ID = uuid.New()

	if err := h.DB.Create(board).Error; err != nil {
		RespondWithError(w, http.StatusInternalServerError, err)
	}

	RespondWithCode(w, http.StatusCreated, board)
}

func (h *Handler) RetrieveBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	board := model.Board{}

	if err := h.DB.Preload("Phases.Methods").Preload("Phases.Phase").First(&board, "id = ?", vars["boardId"]).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		RespondEmptyWithCode(w, http.StatusNotFound)
		return
	} else if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	RespondWithSuccess(w, board)
}

func (h *Handler) DeleteBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) RetrieveBoardGoals(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) CreateGoalInBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) UpdateGoalInBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) RemoveGoalFromBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) RetrieveBoardContents(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) CreateContentFromBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) UpdateContentFromBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) DeleteContentFromBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) RetrieveBoardPhases(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) AddPhaseToBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	phaseId := &model.PhaseIdentifier{}

	if !HandleBodyDecode(w, r, phaseId) {
		return
	}

	boardId, parsingErr := uuid.Parse(vars["boardId"])
	if parsingErr != nil {
		RespondWithCode(w, http.StatusBadRequest, parsingErr)
		return
	}

	if err := h.DB.Model(&model.Board{ID: boardId}).Association("Phases").Append(&model.BoardPhase{ID: uuid.New(), PhaseID: phaseId.ID}); errors.Is(err, gorm.ErrRecordNotFound) {
		RespondEmptyWithCode(w, http.StatusNotFound)
		return
	}

	RespondWithCode(w, http.StatusCreated, phaseId)
}

func (h *Handler) UpdatePhaseInBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) RemovePhaseFromBoard(w http.ResponseWriter, r *http.Request) {
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
