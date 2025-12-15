package main

import (
	"log"
	"r_d/config"
	"r_d/database"

	"github.com/k0kubun/pp/v3"
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

	arr := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five"}
	pp.Println("TEST Go package - PP: ", arr)

	server := NewServer(db)
	server.Run(cfg.Port)
}
