CREATE TABLE IF NOT EXISTS products(
    id VARCHAR(255) PRIMARY KEY, -- Handle UUID Effectively is use Varchar, Optimized Indexing Schema.
    name VARCHAR(255) NOT NULL,
    price DECIMAL,
    is_active BOOLEAN DEFAULT TRUE,
    created_by INT DEFAULT 0, -- Means by System.
    updated_by INT DEFAULT 0, -- Means by System.
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);