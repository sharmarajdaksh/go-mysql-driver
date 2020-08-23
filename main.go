package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	database "github.com/sharmarajdaksh/go-mysql-driver/database"
)

type ordersDBRow struct {
	orderNumber    sql.NullInt32
	orderDate      sql.NullString
	requiredDate   sql.NullString
	shippedDate    sql.NullString
	status         sql.NullString
	comments       sql.NullString
	customerNumber sql.NullInt32
}

func main() {
	mySQLUser := os.Getenv("MYSQL_USER")
	mySQLPassword := os.Getenv("MYSQL_PASSWORD")
	mySQLDBName := os.Getenv("MYSQL_DATABASE")
	dbConnectionString := fmt.Sprintf("%s:%s@/%s", mySQLUser, mySQLPassword, mySQLDBName)

	var err error
	database.DBConnection, err = sql.Open("mysql", dbConnectionString)
	if err != nil {
		panic(err)
	}
	defer database.DBConnection.Close()

	// Configure DB Connection
	database.DBConnection.SetConnMaxLifetime(time.Minute * 3)
	database.DBConnection.SetMaxOpenConns(10)
	database.DBConnection.SetMaxIdleConns(10)

	rows, err := database.DBConnection.Query("SELECT status, customerNumber FROM orders")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var status string
	var customerNumber int
	for rows.Next() {
		if err = rows.Scan(&status, &customerNumber); err != nil {
			panic(err)
		}

		fmt.Println(status, customerNumber)
	}
	if err != nil {
		panic(err)
	}
}
