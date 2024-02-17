package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

type Message struct {
	Sender net.Conn
	Type   string
	Data   string
}

func NewMessage(s net.Conn, t string, d string) *Message {
	return &Message{
		Sender: s,
		Type:   t,
		Data:   d,
	}
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan Message
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		msgch:      make(chan Message, 10),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}

	// defer ln.Close()
	s.ln = ln

	go s.acceptLoop()

	<-s.quitch
	close(s.msgch)

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
			continue
		}

		fmt.Println("new connection to the server: ", conn.RemoteAddr())

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	foundLength := false
	messageLength := 0
	cmdLength := 0
	for {
		if !foundLength {
			var b = make([]byte, 6)
			read, err := conn.Read(b)
			if err != nil {
				if err == io.EOF {
					fmt.Printf("Client %s Disconnected\n", conn.RemoteAddr())
					conn.Close()
					break
				}
				fmt.Println(err)
				continue
			}
			if read != 6 {
				fmt.Print(read)
				fmt.Println("invalid header")
				continue
			}
			foundLength = true
			messageLength = int(binary.LittleEndian.Uint32(b[:4]))
			cmdLength = int(binary.LittleEndian.Uint16(b[4:]))
		} else {
			var command = make([]byte, cmdLength)
			var message = make([]byte, messageLength)
			//Read command
			cmd, err := conn.Read(command)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if cmd != cmdLength {
				fmt.Println("invalid command")
				continue
			}
			//Read Data
			msg, err := conn.Read(message)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if msg != messageLength {
				fmt.Println("invalid message")
				continue
			}

			s.msgch <- Message{
				Sender: conn,
				Type:   string(command),
				Data:   string(message),
			}
			// fmt.Print("Command: ", string(command))
			// fmt.Println("\tData: ", string(message))

			foundLength = false
			messageLength = 0
		}
	}
}

// func handleMessage() {

// }
