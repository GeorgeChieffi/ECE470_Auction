package main

import (
	"fmt"
)

func (c *Client) displayBaseMenu() {
	if c.LoggedInStatus {
		fmt.Println(
			"1 - Logout\n2 - Create an Auction\n3 - List all pending Auctions\n4 - Place a bid\n5 - End an Auction\n6 - Retrieve all Succsesful Bids\n7 - Retrive Winners to your Auctions\n8 - Exit Program")
	} else {
		fmt.Println(
			"1 - Login\n2 - Exit Program")
	}
}

func (c *Client) handleInput() {
	var input int
	fmt.Scanln(&input)

	switch input {
	case 1: // Login
		if c.LoggedInStatus {
			c.LoggedInStatus = false
			break
		}
		c.handleLogin()
	case 2: // Create Auction
		if !c.LoggedInStatus {
			c.exit = true
			break
		}
		c.handleCreateAuction()
	case 3: // List Pending Auctions
		c.handleListPendingAuctions()
	case 4: // Place Bid
		c.handlePlaceBid()
	case 5: // End Auction
		c.handleEndAuction()
	case 6: // Retrieve successful bids
		c.handleRetrieveSuccessfulBids()
	case 7: // Retrieve winners of your auctions
		c.handleRetrieveWinnersOfAuctions()
	case 8: // Exit
		c.exit = true
	}
}
