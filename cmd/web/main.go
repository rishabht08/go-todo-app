package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const portNumber = ":8000"

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbPortNum, err := strconv.Atoi(dbPort)
	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPortNum, dbUser, dbPassword, dbName)

	fmt.Println("Connecting to database ")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// handlers.NewRepo(db)

	fmt.Println("Succesfully Connected to database ")

	fmt.Printf("Staring application on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(db),
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
