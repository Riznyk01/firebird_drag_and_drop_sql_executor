package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/nakagami/firebirdsql"
	"log"
	"sql_executor/internal/config"
)

func NewFirebirdDB(cfg *config.Config, pass, pathToDb string) (*sql.DB, string, error) {

	connectionString := fmt.Sprintf("%s:%s@%s:%s/%s",
		cfg.Login, pass, cfg.Host, cfg.Port, pathToDb)

	log.Printf("connection string is: %s\n", connectionString)
	db, err := sql.Open("firebirdsql", connectionString)

	if err != nil {
		log.Printf("Error connecting to database: %v\n", err)
		return nil, "", err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging database: %v\n", err)
		return nil, "", err
	}
	return db, connectionString, nil
}
