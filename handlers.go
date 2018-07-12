package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func readHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
	}

	// query database for all driver records
	rows, err := driverDB.Query("SELECT * FROM drivers")
	if err != nil {
		panic(err)
	}

	var drivers []Drivers
	var driver Drivers

	// get driver record entry by entry
	for rows.Next() {
		err = rows.Scan(&driver.id, &driver.firstName, &driver.lastName, &driver.age, &driver.team, &driver.entries, &driver.win, &driver.championship)
		if err != nil {
			panic(err)
		}
		drivers = append(drivers, driver)
	}

	fmt.Println("drivers record: ", drivers)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
	}

	// retrieve information about a driver
	var driver Drivers

	driver.id, _ = strconv.Atoi(r.FormValue("id"))
	driver.age, _ = strconv.Atoi(r.FormValue("age"))
	driver.championship, _ = strconv.Atoi(r.FormValue("championship"))
	driver.entries, _ = strconv.Atoi(r.FormValue("entries"))
	driver.firstName = r.FormValue("firstName")
	driver.lastName = r.FormValue("lastName")
	driver.team = r.FormValue("team")
	driver.win, _ = strconv.Atoi(r.FormValue("win"))

	fmt.Println("id = ", driver.id)
	fmt.Println("age = ", driver.age)
	fmt.Println("championship = ", driver.championship)
	fmt.Println("entries = ", driver.entries)
	fmt.Println("firstName = ", driver.firstName)
	fmt.Println("lastName = ", driver.lastName)
	fmt.Println("team = ", driver.team)
	fmt.Println("win = ", driver.win)
	fmt.Println("driver = ", driver)

	// save to database
	prepStatement := `
INSERT INTO drivers(id, firstname, lastname, age, team, entries, win, championship)
VALUES($1, $2, $3, $4, $5, $6, $7, $8)
		`
	// prepStatement := "INSERT INTO drivers(id, age) VALUES($1, $2);"
	sqlStatement, err := driverDB.Prepare(prepStatement)
	if err != nil {
		fmt.Printf("Prepare query error, preped statement = %v, sqlStatement = %v\n", prepStatement, sqlStatement)
		panic(err)
	}

	_, err = sqlStatement.Exec(driver.id, driver.firstName, driver.lastName, driver.age, driver.team, driver.entries, driver.win, driver.championship)
	// _, err = sqlStatement.Exec(driver.id, driver.age)
	if err != nil {
		fmt.Println("Execute query error")
		panic(err)
	}
	http.Redirect(w, r, "/", 301)
}
