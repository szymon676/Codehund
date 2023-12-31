package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/szymon676/codehund/types"
)

func migrateSchemas(db *sql.DB) error {
	userschema := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT UNIQUE,
		email TEXT UNIQUE,
		password TEXT
	);
	`
	postschema := `
	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		author TEXT,
		title TEXT,
		content TEXT
	);
	`
	_, err := db.Exec(userschema)
	if err != nil {
		return err
	}
	_, err = db.Exec(postschema)
	if err != nil {
		return err
	}
	log.Println("migrated user and post schemas")
	return nil
}

func NewPostgresDatabase(connopts *types.PostgresConnectionOptions) *sql.DB {
	connstring := fmt.Sprintf("port=%s user=%s dbname=%s password=%s sslmode=disable", connopts.Port, connopts.User, connopts.DatabaseName, connopts.Password)
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	err = migrateSchemas(db)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
