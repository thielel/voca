package handler

import (
	"encoding/json"
	"net/http"

	"github.com/thielel/voca/internal/domain"
	"github.com/thielel/voca/internal/service"
)

// QuestionnaireHandler handles HTTP requests for the questionnaire
type QuestionnaireHandler struct {
	service *service.PersonalityService
}

// NewQuestionnaireHandler creates a new questionnaire handler
func NewQuestionnaireHandler(svc *service.PersonalityService) *QuestionnaireHandler {
	return &QuestionnaireHandler{service: svc}
}

// GetQuestions handles GET /api/questions
func (h *QuestionnaireHandler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	questions := h.service.GetQuestions()

	response := domain.GetQuestionsResponse{
		Questions: questions,
	}

	writeJSON(w, http.StatusOK, response)
}

// SubmitAnswers handles POST /api/results
func (h *QuestionnaireHandler) SubmitAnswers(w http.ResponseWriter, r *http.Request) {
	var req domain.SubmitAnswersRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if len(req.Answers) == 0 {
		writeError(w, http.StatusBadRequest, "No answers provided")
		return
	}

	result, err := h.service.CalculateResults(req.SessionID, req.Answers, req.Language)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to calculate results")
		return
	}

	response := domain.SubmitAnswersResponse{
		Result: *result,
	}

	writeJSON(w, http.StatusOK, response)
}

// GetResult handles GET /api/results/{id}
func (h *QuestionnaireHandler) GetResult(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		writeError(w, http.StatusBadRequest, "Result ID is required")
		return
	}

	result, err := h.service.GetResult(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to retrieve result")
		return
	}

	if result == nil {
		writeError(w, http.StatusNotFound, "Result not found")
		return
	}

	writeJSON(w, http.StatusOK, result)
}

// GetAllResults handles GET /api/admin/results
func (h *QuestionnaireHandler) GetAllResults(w http.ResponseWriter, r *http.Request) {
	results, err := h.service.GetAllResults()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to retrieve results")
		return
	}

	if results == nil {
		results = []*domain.PersonalityResult{}
	}

	writeJSON(w, http.StatusOK, results)
}

// RegenerateRequest represents the request body for regenerating interpretations
type RegenerateRequest struct {
	Language string `json:"language"`
}

// RegenerateInterpretations handles POST /api/results/{id}/regenerate
func (h *QuestionnaireHandler) RegenerateInterpretations(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		writeError(w, http.StatusBadRequest, "Result ID is required")
		return
	}

	// Parse optional language from request body
	var req RegenerateRequest
	if r.Body != nil {
		_ = json.NewDecoder(r.Body).Decode(&req)
	}

	// Default to German if no language specified
	language := req.Language
	if language == "" {
		language = "de"
	}

	result, err := h.service.RegenerateInterpretations(id, language)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to regenerate interpretations: "+err.Error())
		return
	}

	if result == nil {
		writeError(w, http.StatusNotFound, "Result not found")
		return
	}

	writeJSON(w, http.StatusOK, result)
}

// Helper functions

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
