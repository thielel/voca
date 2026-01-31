package service

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	openai "github.com/sashabaranov/go-openai"
	"github.com/thielel/voca/internal/domain"
)

// OpenAIService handles AI-powered interpretation generation
type OpenAIService struct {
	client *openai.Client
}

// NewOpenAIService creates a new OpenAI service
func NewOpenAIService(apiKey string) *OpenAIService {
	if apiKey == "" {
		return nil
	}
	client := openai.NewClient(apiKey)
	return &OpenAIService{client: client}
}

// GenerateInterpretation creates an AI interpretation for a specific trait and score
func (s *OpenAIService) GenerateInterpretation(ctx context.Context, trait domain.Trait, score float64, language string) (string, error) {
	if s == nil || s.client == nil {
		return "", fmt.Errorf("OpenAI service not configured")
	}

	systemPrompt := GetSystemPrompt(language)
	prompt := BuildInterpretationPrompt(trait, score, language)

	resp, err := s.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: "gpt-5-mini-2025-08-07",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxCompletionTokens: 32000,
		},
	)

	if err != nil {
		return "", fmt.Errorf("failed to generate interpretation: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	content := resp.Choices[0].Message.Content
	log.Printf("OpenAI response for %s (lang=%s): finish_reason=%s, content_length=%d", trait, language, resp.Choices[0].FinishReason, len(content))

	return content, nil
}

// GenerateAllInterpretations generates interpretations for all traits in a result (in parallel)
func (s *OpenAIService) GenerateAllInterpretations(ctx context.Context, result *domain.PersonalityResult, language string) ([]*domain.TraitInterpretation, error) {
	if s == nil || s.client == nil {
		return nil, fmt.Errorf("OpenAI service not configured")
	}

	traits := []struct {
		trait domain.Trait
		score float64
	}{
		{domain.TraitExtraversion, result.Extraversion},
		{domain.TraitAgreeableness, result.Agreeableness},
		{domain.TraitConscientiousness, result.Conscientiousness},
		{domain.TraitEmotionalStability, result.EmotionalStability},
		{domain.TraitOpenness, result.Openness},
	}

	// Pre-allocate slice with fixed positions for thread-safe assignment
	interpretations := make([]*domain.TraitInterpretation, len(traits))

	var wg sync.WaitGroup
	var mu sync.Mutex
	var firstErr error

	for i, t := range traits {
		wg.Add(1)
		go func(idx int, trait domain.Trait, score float64) {
			defer wg.Done()

			interpretation, err := s.GenerateInterpretation(ctx, trait, score, language)
			if err != nil {
				mu.Lock()
				if firstErr == nil {
					firstErr = fmt.Errorf("failed to generate interpretation for %s: %w", trait, err)
				}
				mu.Unlock()
				return
			}

			interpretations[idx] = &domain.TraitInterpretation{
				ID:             uuid.New().String(),
				ResultID:       result.ID,
				Trait:          trait,
				Interpretation: interpretation,
				CreatedAt:      time.Now(),
			}
		}(i, t.trait, t.score)
	}

	wg.Wait()

	if firstErr != nil {
		return nil, firstErr
	}

	return interpretations, nil
}
