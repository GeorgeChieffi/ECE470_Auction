package main

import (
	"fmt"
	"net"
)

type Message struct {
	cmd  string
	data []byte
}

func makeMessage(cmd string, data string) *Message {
	return &Message{
		cmd:  cmd,
		data: make([]byte, len(data)),
	}
}

func main() {
	Message1 := makeMessage("Login", "username:bob@password:pass1")

	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Accept Error: ", err)
	}
	conn.Write(Message1.data)
	conn.Close()
}
