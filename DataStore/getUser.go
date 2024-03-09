package main

import (
	"database/sql"
	"fmt"
)

type User struct {
	UID      int
	Username string
	Password string
	Address  string
}

func getUserByPassword(db *sql.DB, password string) User {
	query := fmt.Sprintf("SELECT * FROM users WHERE Password = '%s';", password)
	var user User

	err := db.QueryRow(query).Scan(&user.UID, &user.Username, &user.Password, &user.Address)
	if err != nil {
		panic(err.Error())
	}
	return user

	// rows, err := db.Query(query)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer rows.Close()

	// // An User slice to hold data from returned rows.
	// var users []User

	// // Loop through rows, using Scan to assign column data to struct fields.
	// for rows.Next() {
	// 	var temp User
	// 	err := rows.Scan(&temp.UID, &temp.Username, &temp.Password, &temp.Address)
	// 	if err != nil {
	// 		return users, err
	// 	}
	// 	users = append(users, temp)
	// }
	// if err = rows.Err(); err != nil {
	// 	return users, err
	// }

	// return users, nil
}
