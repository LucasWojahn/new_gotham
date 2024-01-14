package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDatabase(dotenvPath string) (*sql.DB, error) {
	err := godotenv.Load(dotenvPath)
	if err != nil {
		return nil, err
	}

	//dbUser := os.Getenv("DB_USER")
	//dbPass := os.Getenv("DB_PASS")
	//dbName := os.Getenv("DB_NAME")
	//
	//dbURL := fmt.Sprintf("postgresql://%s:%s@localhost:5432/gotham?sslmode=disable", dbUser, dbPass)

	//connStr := "user=postgres dbname=gotham password= sslmode=disable"

	db, err := sql.Open("postgres", "postgres://postgres:1234@localhost/postgres?sslmode=disable")

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
