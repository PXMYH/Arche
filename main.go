package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

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
	log.Fatal(http.ListenAndServe(":"+port, nil))

	// updating database record
	// 	sqlStatement := `
	// UPDATE users
	// SET age = $2, first_name = $3
	// WHERE id = $1;`
	// 	_, err = db.Exec(sqlStatement, 2, 20, "pos")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Record updated")

	// querying database record
	// type User struct {
	// 	ID        int
	// 	Age       int
	// 	FirstName string
	// 	LastName  string
	// 	Email     string
	// }

	// 	sqlStatement = `
	// SELECT age, first_name, last_name, email FROM users where id = $1;`

	// 	var user_record User
	// 	row := db.QueryRow(sqlStatement, 7)
	// 	err = row.Scan(&user_record.Age, &user_record.FirstName, &user_record.LastName, &user_record.Email)
	// 	switch err {
	// 	case nil:
	// 		fmt.Println("record is retrieved: ", user_record)
	// 	case sql.ErrNoRows:
	// 		fmt.Println("No rows returned, redirecting to 404 now")
	// 	default:
	// 		panic(err)
	// 	}

	// 	sqlStatement = `
	// SELECT id, age, first_name, last_name, email FROM users LIMIT $1;`
	// 	rows, err := db.Query(sqlStatement, 12)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer rows.Close()

	// 	columnNames, err := rows.Columns()
	// 	fmt.Println("names of columns are: ", columnNames)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	columnTypes, err := rows.ColumnTypes()
	// 	fmt.Println("types of columns are: ", columnTypes)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	for rows.Next() {
	// 		err = rows.Scan(&user_record.ID, &user_record.Age, &user_record.FirstName, &user_record.LastName, &user_record.Email)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		fmt.Println("user record => ", user_record)
	// 	}
	// 	err = rows.Err()
	// 	if err != nil {
	// 		panic(err)
	// 	}
}
