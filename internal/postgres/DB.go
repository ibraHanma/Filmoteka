package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Connect() (*sql.DB, error) {
	connStr := "user=postgres password=password dbname=cinematograph host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Ошибка при соединении с базой данных: %v", err)
		return nil, err
	}

	return db, nil
}
