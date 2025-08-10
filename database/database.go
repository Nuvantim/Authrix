package database

import (
	"context"
	"errors"
	"log"
	"sync"

	"api/config"
	"api/internal/app/repository"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DB      *pgxpool.Pool
	once    sync.Once
	Queries *repository.Queries
)

func InitDB() {
	once.Do(func() {
		// Load Database Configuration
		var dbconfig, errs = config.GetDatabaseConfig()
		if errs != nil {
			log.Fatal(errs)
		}
		// Start Connection Database
		var err error
		var dsn string = "postgres://" + dbconfig.User + ":" + dbconfig.Password + "@" + dbconfig.Host + ":" + dbconfig.Port + "/" + dbconfig.Name
		DB, err = pgxpool.New(context.Background(), dsn)
		if err != nil {
			log.Fatalf("Unable to create connection pool: %v", err)
		}

		// Inisialisasi Queries
		Queries = repository.New(DB)
	})
}

func Fatal(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		log.Printf("PGX ERROR | Code: %s | Message: %s | Detail: %s | Where: %s",
			pgErr.Code, pgErr.Message, pgErr.Detail, pgErr.Where)

		return errors.New(pgErr.Message)
	}

	log.Printf("Unexpected error: %v", err)
	return errors.New(err.Error())
}

func CloseDB() {
	log.Println("Disconnection Database")
	if DB != nil {
		defer DB.Close()
	}
}
