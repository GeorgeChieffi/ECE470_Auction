package main

import (
	"fmt"
	"log"
)

func main() {
	server := NewServer(":50000")
	go func() {
		for m := range server.msgch {
			fmt.Printf("[%s] %s:%s\n", m.Sender.RemoteAddr(), m.Type, m.Data)
		}
	}()
	log.Fatal(server.Start())
}
