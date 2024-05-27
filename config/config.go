package config

import "album/internal/database"

type ApiConfig struct {
	DB *database.Queries
}
