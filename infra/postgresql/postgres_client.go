package postgresql

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go-blog-app/config"
	"log"
)

func ConnectPostgres(config config.AppConfig) *sql.DB {
	db, err := sql.Open("postgres", config.DSN)
	if err != nil {
		log.Fatalf("database connection error %v/n", err)
	}
	log.Println("database connection success")
	return db
}
