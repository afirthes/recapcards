package main

import (
	"database/sql"
	"github.com/afirthes/recapcards/internal/db"
	"github.com/afirthes/recapcards/internal/env"
	"github.com/afirthes/recapcards/internal/store"
	"go.uber.org/zap"
)

const version = "0.0.2"

//	@title			Recap Cards API
//	@version		1.0
//	@description	Api for Recap Cards learning app

// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	cfg := config{
		Addr:   env.GetString("SERVER_ADDR", ":8080"),
		ApiURL: env.GetString("EXTERNAL_URL", "localhost:8080"),
		Db: dbConfig{
			addr:         env.GetString("DB_ADDR", ""),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "30s"),
		},
		Env: env.GetString("ENV", "development"),
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer closeLogger(logger)

	// Database
	database, err := db.New(
		cfg.Db.addr,
		cfg.Db.maxOpenConns,
		cfg.Db.maxIdleConns,
		cfg.Db.maxIdleTime)
	if err != nil {
		logger.Fatal(err)
	}

	// Closing database connection
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			logger.Error("Error closing database connection in defer section:", err)
		}
	}(database)

	storage := store.NewPostgresStorage(database)

	app := &application{
		config:  cfg,
		storage: storage,
		logger:  logger,
	}

	logger.Info("Server started at ", "addr", cfg.Addr)
	logger.Fatal(app.Run(app.mount()))
}

func closeLogger(logger *zap.SugaredLogger) {
	err := logger.Sync()
	if err != nil {
		logger.Error("Error syncing logger:", err)
	}
}
