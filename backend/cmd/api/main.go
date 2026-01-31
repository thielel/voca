package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/thielel/voca/internal/config"
	"github.com/thielel/voca/internal/handler"
	"github.com/thielel/voca/internal/repository"
	"github.com/thielel/voca/internal/service"
)

func main() {
	// Load .env file from project root (parent directory)
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := repository.InitDB(cfg.DatabasePath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	resultRepo := repository.NewResultRepository(db)

	// Initialize OpenAI service
	var openaiService *service.OpenAIService
	if cfg.OpenAIAPIKey != "" {
		openaiService = service.NewOpenAIService(cfg.OpenAIAPIKey)
		log.Println("OpenAI service initialized")
	} else {
		log.Println("Warning: OPENAI_API_KEY not set, AI interpretations will be disabled")
	}

	// Initialize services
	personalityService := service.NewPersonalityService(resultRepo, openaiService)

	// Initialize handlers
	questionnaireHandler := handler.NewQuestionnaireHandler(personalityService)

	// Setup routes
	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("GET /api/questions", questionnaireHandler.GetQuestions)
	mux.HandleFunc("POST /api/results", questionnaireHandler.SubmitAnswers)
	mux.HandleFunc("GET /api/results/{id}", questionnaireHandler.GetResult)
	mux.HandleFunc("POST /api/results/{id}/regenerate", questionnaireHandler.RegenerateInterpretations)

	// Admin routes
	mux.HandleFunc("GET /api/admin/results", questionnaireHandler.GetAllResults)

	// Health check
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Apply middleware
	var httpHandler http.Handler = mux
	httpHandler = handler.Logger(httpHandler)
	httpHandler = handler.CORS(httpHandler)

	// Start server
	addr := ":" + cfg.Port
	log.Printf("Starting server on %s", addr)
	log.Printf("Environment: %s", cfg.Environment)
	log.Printf("Database: %s", cfg.DatabasePath)

	if err := http.ListenAndServe(addr, httpHandler); err != nil {
		log.Fatal(err)
	}
}
