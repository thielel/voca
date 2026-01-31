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
func (s *PersonalityService) CalculateResults(sessionID string, answers []domain.Answer) (*domain.PersonalityResult, error) {
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

	// Generate AI interpretations (default to German for initial generation)
	if s.openaiSvc != nil && s.repo != nil {
		ctx := context.Background()
		interpretations, err := s.openaiSvc.GenerateAllInterpretations(ctx, result, "de")
		if err != nil {
			// Log error but don't fail the request - interpretations are optional
			log.Printf("Warning: Failed to generate interpretations: %v", err)
		} else {
			// Save interpretations to database
			if err := s.repo.SaveInterpretations(interpretations); err != nil {
				log.Printf("Warning: Failed to save interpretations: %v", err)
			} else {
				// Add interpretations to result
				result.Interpretations = make(map[domain.Trait]string)
				for _, interp := range interpretations {
					result.Interpretations[interp.Trait] = interp.Interpretation
				}
			}
		}
	}

	return result, nil
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
