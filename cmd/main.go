package main

import (
	"database/sql"
	"log"

	"github.com/mattkibbler/rivers-backend/api"
	"github.com/mattkibbler/rivers-backend/services/tiles"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	addr := ":8080"
	dbPath := "default.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	tilesService := tiles.NewService(db)
	apiServer := api.NewApiServer(addr, db)
	apiServer.RegisterService(tilesService)

	log.Println("Starting server at", addr)
	log.Fatal(apiServer.Start())
}
