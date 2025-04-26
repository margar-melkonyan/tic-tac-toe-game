package config

import (
	"log/slog"

	"github.com/joho/godotenv"
)

func mustLoadEnv() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("can't load .env")
		panic(err)
	}
}
