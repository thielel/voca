package repository

import (
	"database/sql"
	"errors"

	"github.com/thielel/voca/internal/domain"
)

// ErrNotFound is returned when a result is not found
var ErrNotFound = errors.New("result not found")

// ResultRepository handles database operations for personality results
type ResultRepository struct {
	db *sql.DB
}

// NewResultRepository creates a new result repository
func NewResultRepository(db *sql.DB) *ResultRepository {
	return &ResultRepository{db: db}
}

// Save stores a personality result in the database
func (r *ResultRepository) Save(result *domain.PersonalityResult) error {
	query := `
		INSERT INTO personality_results (
			id, session_id, extraversion, agreeableness, 
			conscientiousness, emotional_stability, openness, created_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query,
		result.ID,
		result.SessionID,
		result.Extraversion,
		result.Agreeableness,
		result.Conscientiousness,
		result.EmotionalStability,
		result.Openness,
		result.CreatedAt.Format("2006-01-02 15:04:05"),
	)

	return err
}

// GetByID retrieves a personality result by its ID
func (r *ResultRepository) GetByID(id string) (*domain.PersonalityResult, error) {
	query := `
		SELECT id, session_id, extraversion, agreeableness, 
			conscientiousness, emotional_stability, openness, created_at
		FROM personality_results
		WHERE id = ?
	`

	result := &domain.PersonalityResult{}
	var createdAtStr string
	err := r.db.QueryRow(query, id).Scan(
		&result.ID,
		&result.SessionID,
		&result.Extraversion,
		&result.Agreeableness,
		&result.Conscientiousness,
		&result.EmotionalStability,
		&result.Openness,
		&createdAtStr,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetBySessionID retrieves all results for a session
func (r *ResultRepository) GetBySessionID(sessionID string) ([]*domain.PersonalityResult, error) {
	query := `
		SELECT id, session_id, extraversion, agreeableness, 
			conscientiousness, emotional_stability, openness, created_at
		FROM personality_results
		WHERE session_id = ?
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*domain.PersonalityResult
	for rows.Next() {
		result := &domain.PersonalityResult{}
		var createdAtStr string
		err := rows.Scan(
			&result.ID,
			&result.SessionID,
			&result.Extraversion,
			&result.Agreeableness,
			&result.Conscientiousness,
			&result.EmotionalStability,
			&result.Openness,
			&createdAtStr,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, rows.Err()
}

// GetAll retrieves all personality results, ordered by created_at DESC
func (r *ResultRepository) GetAll() ([]*domain.PersonalityResult, error) {
	query := `
		SELECT id, session_id, extraversion, agreeableness, 
			conscientiousness, emotional_stability, openness, created_at
		FROM personality_results
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*domain.PersonalityResult
	for rows.Next() {
		result := &domain.PersonalityResult{}
		var createdAtStr string
		err := rows.Scan(
			&result.ID,
			&result.SessionID,
			&result.Extraversion,
			&result.Agreeableness,
			&result.Conscientiousness,
			&result.EmotionalStability,
			&result.Openness,
			&createdAtStr,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, rows.Err()
}

// SaveInterpretation stores a trait interpretation in the database
func (r *ResultRepository) SaveInterpretation(interp *domain.TraitInterpretation) error {
	query := `
		INSERT INTO trait_interpretations (
			id, result_id, trait, interpretation, created_at
		) VALUES (?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query,
		interp.ID,
		interp.ResultID,
		string(interp.Trait),
		interp.Interpretation,
		interp.CreatedAt.Format("2006-01-02 15:04:05"),
	)

	return err
}

// SaveInterpretations stores multiple trait interpretations in the database
func (r *ResultRepository) SaveInterpretations(interpretations []*domain.TraitInterpretation) error {
	for _, interp := range interpretations {
		if err := r.SaveInterpretation(interp); err != nil {
			return err
		}
	}
	return nil
}

// GetInterpretationsByResultID retrieves all interpretations for a result
func (r *ResultRepository) GetInterpretationsByResultID(resultID string) (map[domain.Trait]string, error) {
	query := `
		SELECT trait, interpretation
		FROM trait_interpretations
		WHERE result_id = ?
	`

	rows, err := r.db.Query(query, resultID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	interpretations := make(map[domain.Trait]string)
	for rows.Next() {
		var trait string
		var interpretation string
		err := rows.Scan(&trait, &interpretation)
		if err != nil {
			return nil, err
		}
		interpretations[domain.Trait(trait)] = interpretation
	}

	return interpretations, rows.Err()
}

// GetByIDWithInterpretations retrieves a personality result by its ID including interpretations
func (r *ResultRepository) GetByIDWithInterpretations(id string) (*domain.PersonalityResult, error) {
	result, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}

	interpretations, err := r.GetInterpretationsByResultID(id)
	if err != nil {
		return nil, err
	}

	result.Interpretations = interpretations
	return result, nil
}

// DeleteInterpretationsByResultID deletes all interpretations for a result
func (r *ResultRepository) DeleteInterpretationsByResultID(resultID string) error {
	query := `DELETE FROM trait_interpretations WHERE result_id = ?`
	_, err := r.db.Exec(query, resultID)
	return err
}
