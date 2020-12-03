package db

import (
	"database/sql"
	"log"
)

type employeeItem struct {
	ID     int    `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
}

func getconn() (*sql.DB, error) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "priya@1298"
	dbName := "go_account"
	conn, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		log.Println("Something Went Wrong")
		panic(err.Error())
	}
	log.Println("DataBase is Connected")

	return conn, nil

}
