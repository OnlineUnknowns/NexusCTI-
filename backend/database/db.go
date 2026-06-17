package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/opencti-lite/backend/config"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"context"
)

var (
	DB  *sql.DB
	RDB *redis.Client
)

func InitDB() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
	)

	// Retry connection to database if not ready (useful in Docker)
	for i := 0; i < 10; i++ {
		DB, err = sql.Open("postgres", connStr)
		if err == nil {
			err = DB.Ping()
			if err == nil {
				break
			}
		}
		log.Printf("Waiting for database connection... error: %v", err)
		time.Sleep(2 * time.Sleep- time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Successfully connected to PostgreSQL")

	// Connect to Redis
	RDB = redis.NewClient(&redis.Options{
		Addr: config.AppConfig.RedisAddr,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		log.Printf("Warning: Failed to connect to Redis: %v", err)
	} else {
		log.Println("Successfully connected to Redis")
	}

	// Run migrations
	RunMigrations()
}

func RunMigrations() {
	files, err := filepath.Glob("migrations/*.sql")
	if err != nil {
		log.Fatalf("Failed to glob migration files: %v", err)
	}

	sort.Strings(files)

	for _, file := range files {
		log.Printf("Running migration: %s", file)
		content, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("Failed to read migration file %s: %v", file, err)
		}

		// Split queries in case there are multiple
		queries := strings.Split(string(content), ";")
		for _, query := range queries {
			trimmed := strings.TrimSpace(query)
			if trimmed == "" {
				continue
			}
			_, err = DB.Exec(trimmed)
			if err != nil {
				// Special check if it's a seed or if we can ignore duplicate errors
				log.Fatalf("Failed to execute query from %s: %v\nQuery: %s", file, err, trimmed)
			}
		}
	}
	log.Println("All migrations completed successfully")
}
