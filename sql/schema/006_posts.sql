-- +goose Up
CREATE TABLE posts(
    id UUID PRIMARY KEY,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    published_at TIMESTAMP NOT NULL,
    url TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE posts;