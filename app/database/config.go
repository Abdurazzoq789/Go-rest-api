package database

import "fmt"

var (
	dbUsername = "your username"
	dbPassword = "your password"
	dbHost = "your host maybe: localhost"
	dbTable = "your table name"
	dbPort = "psql connected port"
	pgConnStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbTable, dbPassword)
)
