package main

func main() {
	client := NewClient(":50000")
	client.Dial()

	client.sendMessage("LOGIN", "George@PASS")
	client.sendMessage("LOGOUT", "")

	client.conn.Close()
}
