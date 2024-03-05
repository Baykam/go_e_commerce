package main

import (
	"golang_testing_grpc/pkg/config"
	"golang_testing_grpc/pkg/db"

	"github.com/quangdangfit/gocommon/logger"
)

func main() {
	cfg := config.LoadConfig()
	logger.Initialize(cfg.Environment)

	db, err := db.NewDatabase(cfg.DatabaseUri)
	if err != nil {
		logger.Fatal("Cannot Connect to database: %v", err)
	}
	defer db.CloseDatabase()
}
