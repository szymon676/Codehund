package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/szymon676/codehund/types"
)

func NewPostgresDatabase(connopts *types.PostgresConnectionOptions) *sqlx.DB {
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT UNIQUE,
		email TEXT UNIQUE,
		password TEXT
	);
	`
	connstring := fmt.Sprintf("port=%s user=%s dbname=%s password=%s sslmode=disable", connopts.Port, connopts.User, connopts.DatabaseName, connopts.Password)
	db, err := sqlx.Connect("postgres", connstring)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("migrated user schema")
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
