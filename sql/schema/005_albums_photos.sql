-- +goose Up
ALTER TABLE albums
DROP COLUMN photos;

-- +goose Down
ALTER TABLE albums
ADD COLUMN photos BYTEA ARRAY DEFAULT;