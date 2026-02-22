package setup

import (
	"User/internal/repository"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB(dsn string) (repository.UserRepository, error) {

	fmt.Printf("rdebug dsn:%+v", dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Fatal("DB ping failed:", err)
		return nil, err
	}
	log.Println("Connected to Postgres successfully!")
	userRepo := repository.NewUserRepo(db)
	return userRepo, nil
}
