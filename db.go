package main

import (
	"database/sql"
	f "fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Nightpacific@07@tcp(localhost:3306)/Employdb")
	if err != nil {
		f.Println("error")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		f.Print("error with Connecting db.Ping")
		panic(err.Error())
	}
	insert, err := db.Query("INSERT INTO Employee(Eid, Ename, Contact_NO) VALUES (101, 'Prasant',97331);")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	f.Print("Sucessful CONNECTION")
}
