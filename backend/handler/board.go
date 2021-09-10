package handler

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nflow/lesson_organizer/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func rankOrder(db *gorm.DB) *gorm.DB {
	return db.Order("rank")
}

func (h *Handler) RetrieveBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	board := model.Board{}

	if err := h.DB.Preload("Contents", rankOrder).Preload("Phases", rankOrder).Preload("Phases.Methods", rankOrder).Preload("Phases.Methods.Contents", rankOrder).Preload("Phases.Methods.Method").Preload("Phases.Phase").First(&board, "id = ?", vars["boardId"]).Error; errors.Is(err, gorm.ErrRecordNotFound) {
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

func (h *Handler) AddContentToBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	payload := &model.CreateContentDto{}

	if !HandleBodyDecode(w, r, payload) {
		return
	}

	var boardId uuid.UUID
	if !parseUUID(w, vars["boardId"], &boardId) {
		return
	}

	newContent := &model.BoardContent{ID: uuid.New(), BoardID: boardId, BoardMethodID: uuid.Nil, Text: payload.Text, Rank: 0}

	if err := h.DB.Model(&model.Board{ID: boardId}).Association("Contents").Append(newContent); errors.Is(err, gorm.ErrRecordNotFound) {
		RespondEmptyWithCode(w, http.StatusNotFound)
		return
	}

	RespondWithCode(w, http.StatusCreated, newContent)
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

	var lastPhase = &model.BoardPhase{}
	if err := h.DB.Where("board_id = ?", boardId).Order("rank desc").FirstOrInit(lastPhase, &model.BoardPhase{Rank: 0}).Error; err != nil {
		RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	if err := h.DB.Model(&model.Board{ID: boardId}).Association("Phases").Append(&model.BoardPhase{ID: uuid.New(), PhaseID: phaseId.ID, Rank: lastPhase.Rank + 100}); errors.Is(err, gorm.ErrRecordNotFound) {
		RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	RespondWithCode(w, http.StatusCreated, phaseId)
}

func (h *Handler) MovePhaseInBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	moveEntity := &model.MovePhaseDto{}

	if !HandleBodyDecode(w, r, moveEntity) {
		return
	}

	var boardId uuid.UUID
	if !parseUUID(w, vars["boardId"], &boardId) {
		return
	}

	predecessor := model.BoardPhase{}
	if moveEntity.AfterPhaseID != uuid.Nil {
		if err := h.DB.First(&predecessor, "id = ? AND board_id = ?", moveEntity.AfterPhaseID, boardId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			RespondEmptyWithCode(w, http.StatusNotFound)
			return
		} else if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		predecessor.Rank = 0
	}

	successor := model.BoardPhase{}
	if err := h.DB.Order("rank").First(&successor, "board_id = ? AND rank > ?", boardId, predecessor.Rank).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		successor.Rank = predecessor.Rank + 100
	} else if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	rankDelta := (successor.Rank - predecessor.Rank) / 2
	if rankDelta <= 0 {
		var results []model.BoardPhase
		result := h.DB.Model(&model.BoardPhase{}).Order("rank").Where("board_id = ?", boardId).FindInBatches(&results, 100, func(tx *gorm.DB, batch int) error {
			var currentRank uint = 0
			for index := range results {
				currentRank += 100
				results[index].Rank = currentRank

				if results[index].ID == predecessor.ID {
					predecessor = results[index]
				}
				if results[index].ID == successor.ID {
					successor = results[index]
				}
			}
			tx.Save(&results)
			rankDelta = (successor.Rank - predecessor.Rank) / 2

			return nil
		})
		if result.Error != nil {
			RespondWithError(w, http.StatusInternalServerError, result.Error)
			return
		}
	}
	if err := h.DB.Model(&model.BoardPhase{ID: moveEntity.PhaseID}).Update("rank", predecessor.Rank+rankDelta).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		RespondEmptyWithCode(w, http.StatusNotFound)
		return
	} else if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	RespondWithSuccess(w, moveEntity)
}

func (h *Handler) RemovePhaseFromBoard(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) RetrievePhaseMethods(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var phaseId uuid.UUID
	if !parseUUID(w, vars["phaseId"], &phaseId) {
		return
	}

	var methods []model.BoardMethod
	if err := h.DB.Order("rank").Preload(clause.Associations).Find(&methods, "board_phase_id = ?", phaseId).Error; err != nil {
		RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	RespondWithSuccess(w, methods)
}

func (h *Handler) AddMethodToPhase(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	methodId := &model.MethodIdentifier{}

	if !HandleBodyDecode(w, r, methodId) {
		return
	}

	var phaseId uuid.UUID
	if !parseUUID(w, vars["phaseId"], &phaseId) {
		return
	}

	var lastMethod = &model.BoardMethod{}
	if err := h.DB.Where("board_phase_id = ?", phaseId).Order("rank desc").FirstOrInit(lastMethod, &model.BoardMethod{Rank: 0}).Error; err != nil {
		RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	if err := h.DB.Model(&model.BoardPhase{ID: phaseId}).Association("Methods").Append(&model.BoardMethod{ID: uuid.New(), MethodID: methodId.ID, Rank: lastMethod.Rank + 100}); errors.Is(err, gorm.ErrRecordNotFound) {
		RespondEmptyWithCode(w, http.StatusNotFound)
		return
	}

	RespondWithCode(w, http.StatusCreated, phaseId)
}

func (h *Handler) MoveMethod(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	moveEntity := &model.MoveMethodDto{}

	if !HandleBodyDecode(w, r, moveEntity) {
		return
	}

	var phaseId uuid.UUID
	if !parseUUID(w, vars["phaseId"], &phaseId) {
		return
	}

	method := &model.BoardMethod{}
	if err := h.DB.First(method, "id = ? AND board_phase_id = ?", moveEntity.MethodID, phaseId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			RespondEmptyWithCode(w, http.StatusNotFound)
		} else {
			RespondWithError(w, http.StatusInternalServerError, err)
		}
		return
	}

	if moveEntity.ParentPhaseID != uuid.Nil {
		phase := &model.BoardPhase{}
		if err := h.DB.First(phase, "id = ?", moveEntity.ParentPhaseID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				RespondEmptyWithCode(w, http.StatusNotFound)
			} else {
				RespondWithError(w, http.StatusInternalServerError, err)
			}
			return
		}
		method.BoardPhaseID = phase.ID
		phaseId = phase.ID
	}

	predecessor := model.BoardMethod{}
	if moveEntity.AfterMethodID != uuid.Nil {
		if err := h.DB.First(&predecessor, "id = ? AND board_phase_id = ?", moveEntity.AfterMethodID, phaseId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				RespondEmptyWithCode(w, http.StatusNotFound)
			} else {
				RespondWithError(w, http.StatusInternalServerError, err)
			}
			return
		}
	} else {
		predecessor.Rank = 0
	}

	successor := model.BoardMethod{}
	if err := h.DB.Order("rank").First(&successor, "board_phase_id = ? AND rank > ?", phaseId, predecessor.Rank).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			successor.Rank = predecessor.Rank + 100
		} else {
			RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
	}

	rankDelta := (successor.Rank - predecessor.Rank) / 2
	if rankDelta <= 0 {
		var results []model.BoardMethod
		result := h.DB.Model(&model.BoardMethod{}).Order("rank").Where("board_phase_id = ?", phaseId).FindInBatches(&results, 100, func(tx *gorm.DB, batch int) error {
			var currentRank uint = 0
			for index := range results {
				currentRank += 100
				results[index].Rank = currentRank

				if results[index].ID == predecessor.ID {
					predecessor = results[index]
				}
				if results[index].ID == successor.ID {
					successor = results[index]
				}
			}
			tx.Save(&results)
			rankDelta = (successor.Rank - predecessor.Rank) / 2

			return nil
		})
		if result.Error != nil {
			RespondWithError(w, http.StatusInternalServerError, result.Error)
			return
		}
	}

	method.Rank = predecessor.Rank + rankDelta

	if err := h.DB.Save(method).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			RespondEmptyWithCode(w, http.StatusNotFound)
		} else {
			RespondWithError(w, http.StatusInternalServerError, err)
		}
	}

	RespondWithSuccess(w, moveEntity)
}

func DeleteMethodFromPhase(w http.ResponseWriter, r *http.Request) {
	return
}

func RetrieveMethodConents(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) AddContentToMethod(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	payload := &model.CreateContentDto{}

	if !HandleBodyDecode(w, r, payload) {
		return
	}

	var boardId, phaseId, methodId uuid.UUID
	if !parseUUID(w, vars["boardId"], &boardId) || !parseUUID(w, vars["phaseId"], &phaseId) || !parseUUID(w, vars["methodId"], &methodId) {
		return
	}

	newContent := &model.BoardContent{ID: uuid.New(), BoardID: uuid.Nil, BoardMethodID: methodId, Text: payload.Text, Rank: 0}

	if err := h.DB.Model(&model.BoardMethod{ID: methodId}).Association("Contents").Append(newContent); errors.Is(err, gorm.ErrRecordNotFound) {
		RespondEmptyWithCode(w, http.StatusNotFound)
		return
	}

	RespondWithCode(w, http.StatusCreated, newContent)
}

func UpdateContentInMethod(w http.ResponseWriter, r *http.Request) {
	return
}

func RemoveContentFromMethod(w http.ResponseWriter, r *http.Request) {
	return
}
