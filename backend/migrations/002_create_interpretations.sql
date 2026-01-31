-- Create trait_interpretations table for SQLite
CREATE TABLE IF NOT EXISTS trait_interpretations (
    id TEXT PRIMARY KEY,
    result_id TEXT NOT NULL REFERENCES personality_results(id),
    trait TEXT NOT NULL,
    interpretation TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    UNIQUE(result_id, trait)
);

-- Create index on result_id for faster lookups
CREATE INDEX IF NOT EXISTS idx_trait_interpretations_result_id 
ON trait_interpretations(result_id);
