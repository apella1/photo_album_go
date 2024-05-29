-- +goose Up
ALTER TABLE photos
ADD COLUMN img_url TEXT;

-- +goose Down
ALTER TABLE photos
DROP COLUMN img_url;