package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func InitDB() {
	urlExample := "postgres://postgres:P@ssw0rd@localhost:5433/tasks?sslmode=disable"
	var err error
	DB, err = pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping the database: %v\n", err)
		os.Exit(1)
	}

	log.Printf("Connected to the database successfully\n")
}
