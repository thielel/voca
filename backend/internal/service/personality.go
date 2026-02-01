package service

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/thielel/voca/internal/domain"
	"github.com/thielel/voca/internal/repository"
)

// PersonalityService handles personality test business logic
type PersonalityService struct {
	repo      *repository.ResultRepository
	openaiSvc *OpenAIService
}

// NewPersonalityService creates a new personality service
func NewPersonalityService(repo *repository.ResultRepository, openaiSvc *OpenAIService) *PersonalityService {
	return &PersonalityService{
		repo:      repo,
		openaiSvc: openaiSvc,
	}
}

// GetQuestions returns all questionnaire items
func (s *PersonalityService) GetQuestions() []domain.Question {
	return domain.GetQuestions()
}

// CalculateResults processes answers and calculates personality scores
// Interpretations are generated in the background and won't be included in the returned result
func (s *PersonalityService) CalculateResults(sessionID string, answers []domain.Answer, language string) (*domain.PersonalityResult, error) {
	questions := domain.GetQuestions()
	questionMap := make(map[int]domain.Question)
	for _, q := range questions {
		questionMap[q.ID] = q
	}

	// Calculate scores for each trait
	traitScores := make(map[domain.Trait][]float64)
	for _, answer := range answers {
		question, exists := questionMap[answer.QuestionID]
		if !exists {
			continue
		}

		score := float64(answer.Value)
		if question.Reversed {
			score = 6 - score // Reverse the score (1->5, 2->4, 3->3, 4->2, 5->1)
		}

		traitScores[question.Trait] = append(traitScores[question.Trait], score)
	}

	// Calculate normalized scores (0-100)
	calculateNormalizedScore := func(scores []float64) float64 {
		if len(scores) == 0 {
			return 0
		}
		var sum float64
		for _, s := range scores {
			sum += s
		}
		average := sum / float64(len(scores))
		// Normalize from 1-5 scale to 0-100
		return math.Round(((average - 1) / 4) * 100)
	}

	result := &domain.PersonalityResult{
		ID:                 uuid.New().String(),
		SessionID:          sessionID,
		Extraversion:       calculateNormalizedScore(traitScores[domain.TraitExtraversion]),
		Agreeableness:      calculateNormalizedScore(traitScores[domain.TraitAgreeableness]),
		Conscientiousness:  calculateNormalizedScore(traitScores[domain.TraitConscientiousness]),
		EmotionalStability: calculateNormalizedScore(traitScores[domain.TraitEmotionalStability]),
		Openness:           calculateNormalizedScore(traitScores[domain.TraitOpenness]),
		CreatedAt:          time.Now(),
	}

	// Save to repository
	if s.repo != nil {
		if err := s.repo.Save(result); err != nil {
			return nil, err
		}
	}

	// Default language to German if not specified
	if language == "" {
		language = "de"
	}

	// Generate AI interpretations in the background (non-blocking)
	if s.openaiSvc != nil && s.repo != nil {
		go s.generateInterpretationsInBackground(result.ID, result, language)
	}

	return result, nil
}

// Background generation timeout (should be longer than individual call timeouts * retries)
const backgroundGenerationTimeout = 5 * time.Minute

// generateInterpretationsInBackground generates and saves AI interpretations asynchronously
func (s *PersonalityService) generateInterpretationsInBackground(resultID string, result *domain.PersonalityResult, language string) {
	// Panic recovery - don't let a panic crash the background goroutine
	defer func() {
		if r := recover(); r != nil {
			log.Printf("PANIC in background interpretation generation for result %s: %v", resultID, r)
		}
	}()

	startTime := time.Now()
	log.Printf("Starting background interpretation generation for result %s (language: %s)", resultID, language)

	// Create a context with timeout for the entire operation
	ctx, cancel := context.WithTimeout(context.Background(), backgroundGenerationTimeout)
	defer cancel()

	interpretations, err := s.openaiSvc.GenerateAllInterpretations(ctx, result, language)
	if err != nil {
		log.Printf("Warning: Failed to generate interpretations for result %s: %v (elapsed: %v)", resultID, err, time.Since(startTime))
		return
	}

	// Check if we got any interpretations
	if len(interpretations) == 0 {
		log.Printf("Warning: No interpretations generated for result %s (elapsed: %v)", resultID, time.Since(startTime))
		return
	}

	// Save interpretations to database
	if err := s.repo.SaveInterpretations(interpretations); err != nil {
		log.Printf("Warning: Failed to save interpretations for result %s: %v (elapsed: %v)", resultID, err, time.Since(startTime))
		return
	}

	log.Printf("Successfully generated and saved %d interpretations for result %s (elapsed: %v)", len(interpretations), resultID, time.Since(startTime))
}

// GetResult retrieves a personality result by ID (including interpretations)
func (s *PersonalityService) GetResult(id string) (*domain.PersonalityResult, error) {
	if s.repo == nil {
		return nil, nil
	}
	return s.repo.GetByIDWithInterpretations(id)
}

// GetAllResults retrieves all personality results
func (s *PersonalityService) GetAllResults() ([]*domain.PersonalityResult, error) {
	if s.repo == nil {
		return nil, nil
	}
	return s.repo.GetAll()
}

// RegenerateInterpretations regenerates AI interpretations for an existing result
func (s *PersonalityService) RegenerateInterpretations(id string, language string) (*domain.PersonalityResult, error) {
	if s.repo == nil {
		return nil, nil
	}

	// Get the existing result (without interpretations)
	result, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}

	// Check if OpenAI service is available
	if s.openaiSvc == nil {
		return nil, fmt.Errorf("OpenAI service not configured")
	}

	// Delete existing interpretations
	if err := s.repo.DeleteInterpretationsByResultID(id); err != nil {
		return nil, fmt.Errorf("failed to delete existing interpretations: %w", err)
	}

	// Generate new interpretations with the specified language
	ctx := context.Background()
	interpretations, err := s.openaiSvc.GenerateAllInterpretations(ctx, result, language)
	if err != nil {
		return nil, fmt.Errorf("failed to generate interpretations: %w", err)
	}

	// Save new interpretations
	if err := s.repo.SaveInterpretations(interpretations); err != nil {
		return nil, fmt.Errorf("failed to save interpretations: %w", err)
	}

	// Add interpretations to result
	result.Interpretations = make(map[domain.Trait]string)
	for _, interp := range interpretations {
		result.Interpretations[interp.Trait] = interp.Interpretation
	}

	return result, nil
}
