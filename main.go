package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const user = "admin"
const password = "something"
const hostname = "someserver"
const db = "test"

func main() {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, hostname, db))

	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec("gogo", "huh?", "2012-12-09")

	if err != nil {
		fmt.Println(err)
	}

}
