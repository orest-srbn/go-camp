-- +goose Up
CREATE TABLE IF NOT EXISTS articles (
    id SERIAL PRIMARY KEY,
    guid VARCHAR(255) NOT NULL UNIQUE,
    title TEXT NOT NULL,
    link TEXT NOT NULL,
    description TEXT,
    pub_date TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_articles_guid ON articles(guid);
CREATE INDEX IF NOT EXISTS idx_articles_pub_date ON articles(pub_date);

-- +goose Down
DROP TABLE IF EXISTS articles; 