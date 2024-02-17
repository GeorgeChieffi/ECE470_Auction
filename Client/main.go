package main

import (
	"fmt"
	"net"
)

func main() {
	LOGIN_Message := MakeMessage("LOGIN", "username=bob;password=pass1")
	LOGOUT_Message := MakeMessage("LOGOUT", "username=bob")

	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Accept Error: ", err)
	}
	conn.Write(LOGIN_Message.GenerateBinaryMessage())
	conn.Write(LOGOUT_Message.GenerateBinaryMessage())
	conn.Close()
}
