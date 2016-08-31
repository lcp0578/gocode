package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:lcp0578@/go-mysql")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// err = db.Ping()
	// if err != nil {
	// 	panic(err.Error())
	// }

	smsIns, err := db.Prepare("INSERT INTO squarenum VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer smsIns.Close()

	smsOut, err := db.Prepare("SELECT squareNumber FROM squarenum WHERE number = ?")
	if err != nil {
		panic(err.Error())
	}
	defer smsOut.Close()

	for i := 25; i < 35; i++ {
		_, err = smsIns.Exec(i, (i * i))
		if err != nil {
			panic(err.Error())
		}
	}

	var squareNum int
	err = smsOut.QueryRow(13).Scan(&squareNum)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The square number of 13 is: %d", squareNum)

	err = smsOut.QueryRow(15).Scan(&squareNum)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The square number of 15 is: %d", squareNum)
}
