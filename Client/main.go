package main

func main() {
	client := NewClient(":50000")
	client.Dial()

	for {
		client.displayBaseMenu()
		client.handleInput()
		if client.exit {
			break
		}
	}

	// fmt.Println("Sending LOGIN Request ... ")
	// client.sendMessage("LOGIN", "George@PASS")

	// time.Sleep(2 * time.Second)

	// fmt.Println("Sending LOGOUT Request ... ")
	// client.sendMessage("LOGOUT", "")

	client.conn.Close()
}
