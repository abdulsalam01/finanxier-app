CREATE TABLE IF NOT EXISTS users(
    id VARCHAR(255) PRIMARY KEY, -- Handle UUID Effectively is use Varchar, Optimized Indexing Schema.
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(100) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_by INT DEFAULT 0, -- Means by System.
    updated_by INT DEFAULT 0, -- Means by System.
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);