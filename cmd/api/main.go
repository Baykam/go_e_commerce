package main

import (
	"golang_testing_grpc/pkg/config"
	"golang_testing_grpc/pkg/db"

	httpserver "golang_testing_grpc/internal/server/http"

	"github.com/quangdangfit/gocommon/logger"
	"github.com/quangdangfit/gocommon/validation"
)

func main() {
	cfg := config.LoadConfig()
	logger.Initialize(cfg.Environment)

	db, err := db.NewDatabase(cfg.DatabaseUri)
	if err != nil {
		logger.Fatal("Cannot Connect to database: %v", err)
	}
	defer db.CloseDatabase()

	validator := validation.New()

	httpSvr := httpserver.NewServer(validator, db)
	if err := httpSvr.Run(); err != nil {
		logger.Fatal(err)
	}
}
