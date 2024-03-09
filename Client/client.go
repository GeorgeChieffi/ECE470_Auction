package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	dialAddr       string
	conn           net.Conn
	LoggedInStatus bool
	exit           bool
}

func NewClient(dialAddr string) *Client {
	return &Client{
		dialAddr:       dialAddr,
		LoggedInStatus: false,
		exit:           false,
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

	// Handle Response
	// fmt.Println(string(message))
	switch {
	case string(message) == "LOGIN_SUCCESS":
		c.LoggedInStatus = true

	case string(message) == "LOGIN_DENIED":
		c.LoggedInStatus = false

	case len(message) > 6 && string(message)[:7] == "LISTAUC":
		substrings := strings.Split(string(message)[7:], "&")
		// Print the substrings
		for _, s := range substrings {
			fmt.Println(s)
		}
	case string(message) == "PLACEBID_DENIED":
		fmt.Println("There was an error placing your bid")

	case len(message) >= 14 && string(message)[:14] == "GETWINNINGBIDS":
		substrings := strings.Split(string(message)[14:], "&")
		// Print the substrings
		for _, s := range substrings {
			fmt.Println(s)
		}

	case len(message) >= 10 && string(message)[:10] == "GETWINNERS":
		substrings := strings.Split(string(message)[10:], "&")
		// Print the substrings
		for _, s := range substrings {
			fmt.Println(s)
		}

	default:
		fmt.Println("")
	}
}

func (c *Client) sendMessage(command string, data string) {
	c.conn.Write(MakeMessage(command, data).GenerateBinaryMessage())
	c.readMessage(c.conn)
}
