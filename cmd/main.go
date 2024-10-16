package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/mattkibbler/rivers-backend/api"
	"github.com/mattkibbler/rivers-backend/services/tiles"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	addr := flag.String("listenaddr", ":8080", "the server address")
	dbPath := flag.String("dbpath", "default.db", "the location of the sqlite database")
	flag.Parse()

	db, err := sql.Open("sqlite3", *dbPath)
	if err != nil {
		log.Fatal(err)
	}

	tilesService := tiles.NewService(db)
	apiServer := api.NewApiServer(*addr, db)
	apiServer.RegisterService(tilesService)

	log.Println("Starting server at", *addr)
	log.Fatal(apiServer.Start())
}
