-- +goose Up
CREATE TABLE
    photos (
        id UUID PRIMARY KEY,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        title TEXT NOT NULL,
        body BYTEA DEFAULT NULL,
        album_id UUID NOT NULL REFERENCES albums (id) ON DELETE CASCADE,
        user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE
    );

-- +goose Down
DROP TABLE photos;