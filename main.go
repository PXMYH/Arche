package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1q2w3e4r5t"
	dbname   = "formulaone"
)

var driverDB *sql.DB
var err error

func main() {
	fmt.Println("Starting Database")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println("psql info: ", psqlInfo)
	// check the database connection info
	driverDB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer driverDB.Close()

	// ping database to establish a connection
	err = driverDB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database!")

	os.Setenv("PORT", "7777")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// register handlers
	http.HandleFunc("/list", readHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)

	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))
	log.Fatal(http.ListenAndServe(":"+port, router))

	// log.Fatal(http.ListenAndServe(":12345", router))

}
