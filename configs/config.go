package configs

import (
	"context"
	"log"

	"github.com/go-pg/pg/v10"
)

func ConnectDB() *pg.DB {
	options := &pg.Options {
		User: "postgres",
		Password: "password",
		Addr: "localhost:5432",
		Database: "spudify",
	}

	var db *pg.DB = pg.Connect(options)

	// Check if database is running
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		log.Panicf("Failed to connect to database: %s", err)
	}

	log.Printf("Connected to database")

	return db
}