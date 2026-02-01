package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	openai "github.com/sashabaranov/go-openai"
	"github.com/thielel/voca/internal/domain"
)

const (
	// Timeout for individual API calls
	apiCallTimeout = 60 * time.Second
	// Maximum tokens for interpretation (keeps responses concise and fast)
	maxInterpretationTokens = 800
	// Number of retries for failed API calls
	maxRetries = 3
	// Base delay for exponential backoff
	retryBaseDelay = 1 * time.Second
	// Maximum concurrent API calls (all 5 traits in parallel)
	maxConcurrentCalls = 5
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

	// Create a custom HTTP client with TLS configuration
	// This helps with macOS certificate verification issues
	httpClient := &http.Client{
		Timeout: 120 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
			MaxIdleConns:        10,
			IdleConnTimeout:     30 * time.Second,
			DisableCompression:  false,
			DisableKeepAlives:   false,
			ForceAttemptHTTP2:   true,
		},
	}

	config := openai.DefaultConfig(apiKey)
	config.HTTPClient = httpClient

	client := openai.NewClientWithConfig(config)
	return &OpenAIService{client: client}
}

// GenerateInterpretation creates an AI interpretation for a specific trait and score
// Includes retry logic with exponential backoff
func (s *OpenAIService) GenerateInterpretation(ctx context.Context, trait domain.Trait, score float64, language string) (string, error) {
	if s == nil || s.client == nil {
		return "", fmt.Errorf("OpenAI service not configured")
	}

	systemPrompt := GetSystemPrompt(language)
	prompt := BuildInterpretationPrompt(trait, score, language)

	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			// Exponential backoff: 1s, 2s, 4s
			delay := retryBaseDelay * time.Duration(1<<(attempt-1))
			log.Printf("Retrying %s interpretation (attempt %d/%d) after %v", trait, attempt+1, maxRetries, delay)
			
			select {
			case <-ctx.Done():
				return "", ctx.Err()
			case <-time.After(delay):
			}
		}

		// Create a timeout context for this specific call
		callCtx, cancel := context.WithTimeout(ctx, apiCallTimeout)
		
		resp, err := s.client.CreateChatCompletion(
			callCtx,
			openai.ChatCompletionRequest{
				Model: "gpt-4o-mini",
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
				MaxCompletionTokens: maxInterpretationTokens,
			},
		)
		cancel() // Clean up context

		if err != nil {
			lastErr = err
			log.Printf("OpenAI API error for %s (attempt %d): %v", trait, attempt+1, err)
			continue // Retry
		}

		if len(resp.Choices) == 0 {
			lastErr = fmt.Errorf("no response from OpenAI")
			continue // Retry
		}

		content := resp.Choices[0].Message.Content
		log.Printf("OpenAI response for %s (lang=%s): finish_reason=%s, content_length=%d, attempt=%d",
			trait, language, resp.Choices[0].FinishReason, len(content), attempt+1)

		return content, nil
	}

	return "", fmt.Errorf("failed to generate interpretation after %d attempts: %w", maxRetries, lastErr)
}

// GenerateAllInterpretations generates interpretations for all traits in a result
// Uses bounded parallelism and saves partial results on failure
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
	errors := make([]error, len(traits))

	var wg sync.WaitGroup
	// Semaphore to limit concurrent API calls
	semaphore := make(chan struct{}, maxConcurrentCalls)

	for i, t := range traits {
		wg.Add(1)
		go func(idx int, trait domain.Trait, score float64) {
			defer wg.Done()
			defer func() {
				// Panic recovery
				if r := recover(); r != nil {
					log.Printf("Panic recovered in interpretation generation for %s: %v", trait, r)
					errors[idx] = fmt.Errorf("panic: %v", r)
				}
			}()

			// Acquire semaphore slot
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			interpretation, err := s.GenerateInterpretation(ctx, trait, score, language)
			if err != nil {
				errors[idx] = err
				log.Printf("Failed to generate interpretation for %s: %v", trait, err)
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

	// Collect successful interpretations (partial results are OK)
	var successful []*domain.TraitInterpretation
	var failedTraits []string
	for i, interp := range interpretations {
		if interp != nil {
			successful = append(successful, interp)
		} else if errors[i] != nil {
			failedTraits = append(failedTraits, string(traits[i].trait))
		}
	}

	// Log summary
	log.Printf("Interpretation generation complete: %d/%d successful", len(successful), len(traits))
	if len(failedTraits) > 0 {
		log.Printf("Failed traits: %v", failedTraits)
	}

	// Return partial results if we have any
	if len(successful) > 0 {
		return successful, nil
	}

	// All failed - return the first error
	for _, err := range errors {
		if err != nil {
			return nil, err
		}
	}

	return nil, fmt.Errorf("all interpretation generations failed")
}
