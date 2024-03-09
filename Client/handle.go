package main

import (
	"fmt"
)

func (c *Client) handleLogin() {
	var Username string
	var Password string

	fmt.Print("Enter your Username: ")
	fmt.Scanln(&Username)
	fmt.Print("Enter your Password: ")
	fmt.Scanln(&Password)

	msg := fmt.Sprintf("%s&%s", Username, Password)

	c.sendMessage("LOGIN", msg)
}

func (c *Client) handleCreateAuction() {
	var pName string
	var year string
	var month string
	var day string
	var hour string
	var minute string
	var price string
	fmt.Print("Enter the product name: ")
	fmt.Scanln(&pName)
	fmt.Println("Below enter the ending date.\tExample:\n\tYear: 2024\n\tMonth: 3\n\tDay: 9\n\tHour: 11\n\tMinute: 30")
	fmt.Print("Year: ")
	fmt.Scanln(&year)
	fmt.Print("Month[1-12]: ")
	fmt.Scanln(&month)
	fmt.Print("Day: ")
	fmt.Scanln(&day)
	fmt.Print("Hour[0-23]: ")
	fmt.Scanln(&hour)
	fmt.Print("Minute[0-60]: ")
	fmt.Scanln(&minute)
	fmt.Print("Starting Price: ")
	fmt.Scanln(&price)

	msg := fmt.Sprintf("%s&%s&%s&%s&%s&%s&%s", pName, year, month, day, hour, minute, price)
	c.sendMessage("CREATEAUC", msg)
}

func (c *Client) handleListPendingAuctions() {
	c.sendMessage("LISTAUC", "")
}

func (c *Client) handlePlaceBid() {
	var Aid string
	var amount string

	fmt.Print("Which Auction would you like to bid on(AuctionID): ")
	fmt.Scanln(&Aid)
	fmt.Print("Amount(Must be greater than current price): ")
	fmt.Scanln(&amount)

	msg := fmt.Sprintf("%s&%s", Aid, amount)
	c.sendMessage("PLACEBID", msg)
}

func (c *Client) handleEndAuction() {
	var Aid string
	fmt.Print("Which auction would you like to end: ")
	fmt.Scanln(&Aid)

	c.sendMessage("ENDAUC", Aid)
}

func (c *Client) handleRetrieveSuccessfulBids() {
	c.sendMessage("RETRIEVEBID", "")
}

func (c *Client) handleRetrieveWinnersOfAuctions() {
	c.sendMessage("RETRIEVEWINNERS", "")
}
