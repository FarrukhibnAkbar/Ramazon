package utils

import (
	"fmt"

	"ramazon/configs"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(kind string) (string, error) {
	var config = configs.Load()
	// Define URL to connection.
	var url string

	// Switch given names.
	switch kind {
	case "postgres":
		// URL for PostgreSQL connection.
		url = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.PostgresHost,
			config.PostgresPort,
			config.PostgresUser,
			config.PostgresPassword,
			config.PostgresDatabase,
		)
	case "migration":
		// URL for Migration
		url = fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=disable",
			config.PostgresUser,
			config.PostgresPassword,
			config.PostgresHost,
			config.PostgresPort,
			config.PostgresDatabase,
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", kind)
	}

	// Return connection URL.
	return url, nil
}
