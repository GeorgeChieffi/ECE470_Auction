package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	UID      int
	Username string
	Password string
	Address  string
}

func getUser(username string, password string) (User, error) {
	// connect to SQL Server
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ECE470DataStore")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//Create Query
	query := fmt.Sprintf("SELECT * FROM users WHERE Username COLLATE utf8mb4_bin = '%s' AND Password COLLATE utf8mb4_bin = '%s';", username, password)
	var user User

	err2 := db.QueryRow(query).Scan(&user.UID, &user.Username, &user.Password, &user.Address)
	if err2 != nil {
		if err2 == sql.ErrNoRows {
			// Handle case where no matching rows were found
			return user, errors.New("NOT A VALID USER")
		} else {
			panic(err2.Error())
		}
	}

	return user, nil
}

func getAucs() (string, error) {
	// connect to SQL Server
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ECE470DataStore")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//Create Query
	query := "SELECT a.AuctionID, a.ProductName, a.EndTime, a.StartPrice, MAX(b.Amount) FROM auctions a LEFT JOIN bids b on a.AuctionID = b.AuctionID WHERE a.Closed IS FALSE GROUP BY a.AuctionID;"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	defer rows.Close()

	var allRowsData = "LISTAUC"

	// Iterate over the rows
	for rows.Next() {
		var AuctionID int
		var ProductName string
		var EndTime string
		var StartPrice sql.NullFloat64
		var CurrentBidAmount sql.NullFloat64

		// Scan the columns of the current row into variables
		err := rows.Scan(&AuctionID, &ProductName, &EndTime, &StartPrice, &CurrentBidAmount)
		if err != nil {
			return "", err
		}
		if !CurrentBidAmount.Valid {
			CurrentBidAmount = StartPrice
		}
		// Append the row data as a string to the slice
		rowData := fmt.Sprintf("AuctionID: %d, ProductName: %s, EndTime: %s, Start Price: %.2f, Current Price: %.2f&", AuctionID, ProductName, EndTime, StartPrice.Float64, CurrentBidAmount.Float64)
		allRowsData = allRowsData + rowData
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return "", err
	}
	return allRowsData, nil
}

func insertAuc(pName string, year string, month string, day string, hour string, minute string, price float64, uid int) error {
	// connect to SQL Server
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ECE470DataStore")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := fmt.Sprintf("INSERT INTO auctions (`ProductName`, `EndTime`, `StartPrice`, `Closed`, `Creator`) VALUES ('%s', '%s-%s-%s %s:%s:00', '%f', '%d', '%d')", pName, year, month, day, hour, minute, price, 0, uid)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func createBid(Aid string, amount string, uid int) error {
	// connect to SQL Server
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ECE470DataStore")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query1 := fmt.Sprintf("INSERT INTO bids (`AuctionID`,`UID`,`Amount`) VALUES ('%s','%d','%s');", Aid, uid, amount)

	_, err1 := db.Exec(query1)
	if err1 != nil {
		return err1
	}
	return nil
}

func endAucDB(deletor int, Aid string) error {
	// connect to SQL Server
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ECE470DataStore")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// End Auction
	query1 := fmt.Sprintf("UPDATE auctions SET `Closed` = true WHERE Creator = %d AND `AuctionID` = %s;", deletor, Aid)
	// Update Winner
	query2 := fmt.Sprintf("UPDATE auctions SET `Winner` = (SELECT UID FROM bids WHERE AuctionID = %s GROUP BY UID ORDER BY MAX(Amount) DESC LIMIT 1) WHERE `AuctionID` = %s;", Aid, Aid)

	_, err1 := db.Exec(query1)
	if err1 != nil {
		return err1
	}
	_, err2 := db.Exec(query2)
	if err2 != nil {
		return err2
	}

	return nil
}

func getWinningBids(uid int) (string, error) {
	// connect to SQL Server
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ECE470DataStore")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//Create Query
	query := fmt.Sprintf("SELECT a.AuctionID, a.ProductName, MAX(b.Amount) FROM auctions a LEFT JOIN bids b on a.AuctionID = b.AuctionID WHERE a.Closed IS TRUE AND a.Winner = %d GROUP BY a.AuctionID;", uid)

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	defer rows.Close()

	var allRowsData = "GETWINNINGBIDS"

	// Iterate over the rows
	for rows.Next() {
		var AuctionID int
		var ProductName string
		var FinalBidAmount sql.NullFloat64

		// Scan the columns of the current row into variables
		err := rows.Scan(&AuctionID, &ProductName, &FinalBidAmount)
		if err != nil {
			return "", err
		}
		// Append the row data as a string to the slice
		if FinalBidAmount.Valid {
			rowData := fmt.Sprintf("AuctionID: %d, ProductName: %s, Final Bid Price: %.2f&", AuctionID, ProductName, FinalBidAmount.Float64)
			allRowsData = allRowsData + rowData
		}
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return "", err
	}
	return allRowsData, nil
}

func getWinners(uid int) (string, error) {
	// connect to SQL Server
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ECE470DataStore")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT u.Username, a.Productname FROM auctions a JOIN users u on a.winner = u.UID WHERE a.Creator = %d;", uid)

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	defer rows.Close()

	var allRowsData = "GETWINNERS"

	// Iterate over the rows
	for rows.Next() {
		var Username string
		var ProductName string

		// Scan the columns of the current row into variables
		err := rows.Scan(&Username, &ProductName)
		if err != nil {
			return "", err
		}
		// Append the row data as a string to the slice
		rowData := fmt.Sprintf("Username: %s, ProductName: %s&", Username, ProductName)
		allRowsData = allRowsData + rowData
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return "", err
	}
	return allRowsData, nil
}
