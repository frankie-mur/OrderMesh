package main

import (
	"log"
	"os"

	"github.com/frankie-mur/OrderMesh/api/server"
	"github.com/frankie-mur/OrderMesh/db"
	"github.com/frankie-mur/OrderMesh/ecomm-api/storer"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading the .env file")
	}
	dbConn := os.Getenv("DB_CONN_URL")

	db, err := db.NewDatabase(dbConn)
	if err != nil {
		log.Fatalf("error opening the database: %v", err)
	}
	defer db.Close()
	log.Println("successfully connected to database")

	store := storer.NewPostgresStorer(db.GetDB())
	_ = server.NewServer(store)

}
