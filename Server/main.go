package main

import (
	"log"
)

func main() {
	server := NewServer(":50000")

	// Add the creator of an auction cannot bet on their own auction

	log.Fatal(server.Start())
}
