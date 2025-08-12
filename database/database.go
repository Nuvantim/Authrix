package database

import (
	"context"
	"errors"
	"log"
	"os"
	"sync"
	"time"

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
		dbconfig, errs := config.GetDatabaseConfig()
		if errs != nil {
			log.Fatal(errs)
		}

		dsn := "postgres://" + dbconfig.User + ":" + dbconfig.Password + "@" + dbconfig.Host + ":" + dbconfig.Port + "/" + dbconfig.Name

		// Parse config supaya bisa custom pool
		poolConfig, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			log.Fatalf("Unable to parse database config: %v", err)
		}

		// Pool tuning
		poolConfig.MaxConns = 20                    // batas koneksi maksimum
		poolConfig.MinConns = 5                      // koneksi standby
		poolConfig.MaxConnIdleTime = 5 * time.Minute // idle max 5 menit
		poolConfig.MaxConnLifetime = time.Hour       // koneksi maksimal 1 jam
		poolConfig.HealthCheckPeriod = time.Minute   // cek koneksi tiap menit

		// Statement cache biar query sering dipakai jadi lebih cepat
		poolConfig.ConnConfig.DefaultQueryExecMode = pgxpool.QueryExecModeCacheStatement

		// Timeout default (misal 5 detik)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		DB, err = pgxpool.NewWithConfig(ctx, poolConfig)
		if err != nil {
			log.Fatalf("Unable to create connection pool: %v", err)
		}

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
		DB.Close()
	}
}
