package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1q2w3e4r5t"
	dbname   = "arche"
)

func main() {
	fmt.Println("Starting Database")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println("psql info: ", psqlInfo)
	// check the database connection info
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// ping database to establish a connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	// retrieving database record
	// 	sqlStatement := `
	// INSERT INTO users (age, first_name, last_name, email)
	// VALUES ($1, $2, $3, $4)
	// RETURNING id`
	// id := 0
	// 	err = db.QueryRow(sqlStatement, 24, "Mike", "Hu", "mickal.hu@top.com").Scan(&id)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("New record inserted: ", id)

	// updating database record
	sqlStatement := `
UPDATE users
SET age = $2, first_name = $3
WHERE id = $1;`
	_, err = db.Exec(sqlStatement, 2, 20, "pos")
	if err != nil {
		panic(err)
	}
	fmt.Println("Record updated")

	// querying database record
	type User struct {
		Age       int
		FirstName string
		LastName  string
		Email     string
	}

	sqlStatement = `
SELECT age, first_name, last_name, email FROM users where id = $1;`

	var user_record User
	row := db.QueryRow(sqlStatement, 7)
	err = row.Scan(&user_record.Age, &user_record.FirstName, &user_record.LastName, &user_record.Email)
	switch err {
	case nil:
		fmt.Println("record is retrieved: ", user_record)
	case sql.ErrNoRows:
		fmt.Println("No rows returned, redirecting to 404 now")
	default:
		panic(err)
	}
}
