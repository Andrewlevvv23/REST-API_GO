package main

import (
	"log"
	"r_d/config"
	"r_d/database"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DSN)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer db.Close()

	err = database.CreateTables(db)
	if err != nil {
		log.Fatal("Failed to create tables:", err)
	}

	server := NewServer(db)
	server.Run(cfg.Port)
}
