package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:lcp0578@tcp(127.0.0.1:3306)/go_test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	var result sql.Result
	// INSERT DATA
	result, err = db.Exec("INSERT INTO users(name, age, sex) VALUES (?,?,?)", "lcpeng", 25, 0)
	if err != nil {
		fmt.Println(err)
	}
	// GET　LAST INSERT ID
	lastId, _ := result.LastInsertId()
	fmt.Println("last ID:", lastId)
	// GET ONE ROW
	var row *sql.Row
	row = db.QueryRow("SELECT * FROM users")
	var (
		name         string
		id, age, sex int
	)
	// GET VLAUES
	err = row.Scan(&id, &name, &age, &sex)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id, "\t", name, "\t", age, "\t", sex)
	fmt.Println("============")
	// GET LIST
	var rows *sql.Rows
	rows, err = db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		rows.Scan(&id, &name, &age, &sex)
		fmt.Println(id, "\t", name, "\t", age, "\t", sex)
	}
	rows.Close()
	//
	db.Exec("TRUNCATE TABLE users")
}
