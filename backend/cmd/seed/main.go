package main

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/thielel/voca/internal/config"
	"github.com/thielel/voca/internal/domain"
	"github.com/thielel/voca/internal/repository"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := repository.InitDB(cfg.DatabasePath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize repository
	resultRepo := repository.NewResultRepository(db)

	// Define 5 exemplary personality profiles
	seedData := []struct {
		name               string
		extraversion       float64
		agreeableness      float64
		conscientiousness  float64
		emotionalStability float64
		openness           float64
	}{
		{
			name:               "High Extravert",
			extraversion:       85,
			agreeableness:      60,
			conscientiousness:  55,
			emotionalStability: 50,
			openness:           70,
		},
		{
			name:               "Conscientious Achiever",
			extraversion:       45,
			agreeableness:      70,
			conscientiousness:  90,
			emotionalStability: 65,
			openness:           55,
		},
		{
			name:               "Creative Thinker",
			extraversion:       60,
			agreeableness:      55,
			conscientiousness:  40,
			emotionalStability: 50,
			openness:           95,
		},
		{
			name:               "Agreeable Helper",
			extraversion:       50,
			agreeableness:      92,
			conscientiousness:  60,
			emotionalStability: 70,
			openness:           50,
		},
		{
			name:               "Balanced Individual",
			extraversion:       55,
			agreeableness:      55,
			conscientiousness:  55,
			emotionalStability: 55,
			openness:           55,
		},
	}

	log.Println("Seeding database with 5 exemplary personality results...")

	for i, data := range seedData {
		result := &domain.PersonalityResult{
			ID:                 uuid.New().String(),
			SessionID:          "seed-session-" + uuid.New().String()[:8],
			Extraversion:       data.extraversion,
			Agreeableness:      data.agreeableness,
			Conscientiousness:  data.conscientiousness,
			EmotionalStability: data.emotionalStability,
			Openness:           data.openness,
			CreatedAt:          time.Now().Add(-time.Duration(i) * time.Hour), // Stagger creation times
		}

		if err := resultRepo.Save(result); err != nil {
			log.Printf("Failed to save result for %s: %v", data.name, err)
			continue
		}

		log.Printf("Created: %s (ID: %s)", data.name, result.ID[:8])
		log.Printf("  E: %.0f%%, A: %.0f%%, C: %.0f%%, ES: %.0f%%, O: %.0f%%",
			data.extraversion, data.agreeableness, data.conscientiousness,
			data.emotionalStability, data.openness)
	}

	log.Println("Seeding complete!")
}
