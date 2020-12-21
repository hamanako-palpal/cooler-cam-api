package config

import (
	"database/sql"
	"fmt"
	"os"
)

// CreateDB コネクション作成
func CreateDB() *sql.DB {
	dburi := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_INST"),
		os.Getenv("DB_SSLMODE"))

	db, err := sql.Open("postgres", dburi)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
