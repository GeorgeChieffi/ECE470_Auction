package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

type Client struct {
	dialAddr string
	conn     net.Conn
}

func NewClient(dialAddr string) *Client {
	return &Client{
		dialAddr: dialAddr,
	}
}

func (c *Client) Dial() error {
	conn, err := net.Dial("tcp", c.dialAddr)
	if err != nil {
		fmt.Println("Accept Error: ", err)
	}
	c.conn = conn

	return nil
}

func (c *Client) readMessage(conn net.Conn) {

	//Get Message Length
	var b = make([]byte, 4)
	read, _ := conn.Read(b)

	if read != 4 {
		fmt.Print("Recieved ", read)
		fmt.Println(": Invalid header")
	}

	messageLength := int(binary.LittleEndian.Uint32(b))

	//Read Message
	var message = make([]byte, messageLength)

	msg, err := conn.Read(message)
	if err != nil {
		fmt.Println(err)
	}
	if msg != messageLength {
		fmt.Println("invalid message")
	}

	// Print Response
	fmt.Println(string(message))
}

func (c *Client) sendMessage(command string, data string) {
	c.conn.Write(MakeMessage(command, data).GenerateBinaryMessage())
	c.readMessage(c.conn)
}
