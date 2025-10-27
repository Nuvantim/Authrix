package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"api/config"
	"api/internal/app/repository"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DB      *pgxpool.Pool
	Queries *repository.Queries
	once    sync.Once
)

func InitDB() {
	once.Do(func() {
		dbConfig, err := config.GetDatabaseConfig()
		if err != nil {
			log.Fatalf("Failed to get database config: %v", err)
		}

		dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

		poolConfig, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			log.Fatalf("Unable to parse database config DSN: %v", err)
		}

		poolConfig.MaxConns = 20
		poolConfig.MinConns = 5
		poolConfig.MaxConnIdleTime = 5 * time.Minute
		poolConfig.MaxConnLifetime = time.Hour
		poolConfig.HealthCheckPeriod = time.Minute

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		DB, err = pgxpool.NewWithConfig(ctx, poolConfig)
		if err != nil {
			log.Fatalf("Unable to create connection pool: %v", err)
		}

		if err = DB.Ping(ctx); err != nil {
			log.Fatalf("Could not ping database: %v", err)
		}

		Queries = repository.New(DB)
		fmt.Println("Database connected...")
	})
}

func Fatal(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		log.Printf("PGX ERROR | Code: %s | Message: %s | Detail: %s | Where: %s",
			pgErr.Code, pgErr.Message, pgErr.Detail, pgErr.Where)
		return fmt.Errorf("database error: %s", pgErr.Message)
	}

	log.Printf("Unexpected error: %v", err)
	return err
}

func CloseDB() {
	if DB != nil {
		log.Println("Closing database connection pool.")
		DB.Close()
	}
}
