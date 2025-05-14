package main

import (
	"os"
	"strconv"
)

type config struct {
	PORT                     string
	MONOLITH_URL             string
	MOVIES_SERVICE_URL       string
	EVENTS_SERVICE_URL       string
	GRADUAL_MIGRATION        bool
	MOVIES_MIGRATION_PERCENT int
}

func readCfg() config {
	percent, _ := strconv.Atoi(os.Getenv("MOVIES_MIGRATION_PERCENT"))

	return config{
		PORT:                     os.Getenv("PORT"),
		MONOLITH_URL:             os.Getenv("MONOLITH_URL"),
		MOVIES_SERVICE_URL:       os.Getenv("MOVIES_SERVICE_URL"),
		EVENTS_SERVICE_URL:       os.Getenv("EVENTS_SERVICE_URL"),
		GRADUAL_MIGRATION:        os.Getenv("GRADUAL_MIGRATION") == "true",
		MOVIES_MIGRATION_PERCENT: percent,
	}
}
