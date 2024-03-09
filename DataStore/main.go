package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ECE470DataStore")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	// insert, err := db.Query("INSERT INTO `ECE470DataStore`.`users` (`Username`, `Password`, `Address`) VALUES ('John', 'Doe', '123 Main St.');")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()

	result := getUserByPassword(db, "Doe")
	fmt.Println(result.Username)

	// fmt.Println("Success!")
}
