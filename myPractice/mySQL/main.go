package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:poeticV#1@tcp(127.0.0.1:3306)/hosting_db")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO request VALUES (null, 'AL AMIN', 'MasterAcademy', 'alamincsecu@gmail.com', 1)")

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("connection succrssfull")

}
