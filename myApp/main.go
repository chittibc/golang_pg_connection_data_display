package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

//employee data
type employee struct {
	fName string
	lName string
}

//main function
func main() {
	//Variables required for setup
	/*
		user= (using default user for postgres database)
		dbname= (using default database that comes with postgres)
		password = (password used during initial setup)
		host = (IP Address of server) mine is on my own computer so it's localhost
		sslmode = (must be set to disabled unless using SSL. This is not covered during tutorial)
	*/

	//DO NOT SAVE PASSWORD AS TEXT IN A PRODUCTION ENVIRONMENT. TRY USING AN ENVIRONMENT VARIABLE
	connStr := "user=postgres dbname=postgres password=admin host=localhost sslmode=disable"
	//driver name part of "github.com/lib/pq"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print(err)
	}
	//check postgres to see if table exists
	var checkDatabase string
	db.QueryRow("SELECT to_regclass('public.userList')").Scan(&checkDatabase)
	if err != nil {
		fmt.Print(err)
	}
	//if table dose not exist then create one to use for this example
	if checkDatabase == "" {
		fmt.Println("Database Created")
		createSQL := "CREATE TABLE public.userList (pk SERIAL PRIMARY KEY,fname character varying,lname character varying);"
		db.Query(createSQL)
	}

	//sql to insert employee information
	statement := "INSERT INTO userList(fname, lname) VALUES($1, $2)"
	//prepare statement for sql
	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Print(err)
	}
	defer stmt.Close()
	//call a instant of employee
	eName := employee{}
	//allow 3 employee to be entered into database
	for i := 0; i < 10; i++ {
		fmt.Print("First Name: ")
		//set fName of strut with text input
		fmt.Scanf("%s", &eName.fName)
		fmt.Print("Last Name: ")
		//set fName of strut with text input
		fmt.Scanf("%s", &eName.lName)
		//call prepared statement above
		stmt.QueryRow(eName.fName, eName.lName)
	}
	//select employee first and last name
	rows, err := db.Query("Select fname, lname from userList")
	if err != nil {
		fmt.Print(err)
	}
	defer rows.Close()

	fmt.Println("---------------------------------------------------------------------")
	//loop through all employee results
	for rows.Next() {
		//assign values to variables
		var fname string
		var lname string
		err := rows.Scan(&fname, &lname)
		if err != nil {
			fmt.Print(err)
		}
		//print results to console
		fmt.Printf("%s %s\n", fname, lname)
	} //end of for loop
} //end of main function
