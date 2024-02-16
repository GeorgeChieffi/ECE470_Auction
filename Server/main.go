package main

import (
	"log"
)

func main() {
	server := NewServer(":50000")

	log.Fatal(server.Start())
}
