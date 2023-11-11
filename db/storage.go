package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Init Base on the default connection create a connection with the database name specific for our app
func Init() (*sql.DB, error) {
	/* Load Env */
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading env")
	}

	defaultDbName := "postgres"
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	// Construct the connection string
	defaultDbURL := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbHost, dbPort, defaultDbName)
	dbUrl := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	return setupPosgres(defaultDbURL, dbUrl)
}

func setupPosgres(defaultDbURL, dbUrl string) (*sql.DB, error) {
	/*  Open a connection to the postgres database */
	// sql.Open is used to establish a connection to the PostgreSQL database.
	// However, the sql.Open function only creates a connection object, it doesn't actually establish a connection to the database.
	// In order to use this we also need to import "github.com/lib/pq"
	sqlConnection, dbConnErr := sql.Open("postgres", dbUrl)
	if dbConnErr != nil {
		return nil, dbConnErr
	}

	// sql.Open successfully returns an instance of sql.DB regardless of whether the database server is running or not.
	// To check if the connection was successful, you need to call the Ping method on the sql.DB instance.
	if err := sqlConnection.Ping(); err != nil {
		return nil, createRequireDatabase(defaultDbURL)
	}

	return sqlConnection, nil
}

func createRequireDatabase(defaultDbUrl string) error {
	dbConnection, dbConnErr := sql.Open("postgres", defaultDbUrl)
	if dbConnErr != nil {
		return dbConnErr
	}

	dbName := os.Getenv("POSTGRES_DB")
	query := fmt.Sprintf("CREATE DATABASE %s", dbName)
	_, err := dbConnection.Exec(query)
	return err
}
