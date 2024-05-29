-- +goose Up
ALTER TABLE photos
ALTER COLUMN img_url
SET
    NOT NULL;

-- +goose Down
-- This migration cannot be reverted as making the column non-nullable may result in data loss.