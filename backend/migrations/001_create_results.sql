-- Create personality_results table for SQLite
CREATE TABLE IF NOT EXISTS personality_results (
    id TEXT PRIMARY KEY,
    session_id TEXT NOT NULL,
    extraversion REAL NOT NULL,
    agreeableness REAL NOT NULL,
    conscientiousness REAL NOT NULL,
    emotional_stability REAL NOT NULL,
    openness REAL NOT NULL,
    created_at TEXT NOT NULL DEFAULT (datetime('now'))
);

-- Create index on session_id for faster lookups
CREATE INDEX IF NOT EXISTS idx_personality_results_session_id 
ON personality_results(session_id);

-- Create index on created_at for sorting
CREATE INDEX IF NOT EXISTS idx_personality_results_created_at 
ON personality_results(created_at DESC);
