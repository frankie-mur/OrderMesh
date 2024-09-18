package main

import (
	"context"
	"log"
	"os"

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

	p := storer.Product{
		Name: "Frank",
	}
	newP, err := store.CreateProduct(context.Background(), &p)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Print(newP)

	allP, err := store.ListProducts(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Print(allP)

}
