package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

type MyConn struct {
	conn           net.Conn
	loggedInStatus bool
	msgch          chan Message
	user           User
}

func newConnection(conn net.Conn, loggedInStatus bool) *MyConn {
	return &MyConn{
		conn:           conn,
		loggedInStatus: loggedInStatus,
		msgch:          make(chan Message, 10),
	}
}

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
	listenAddr     string
	ln             net.Listener
	quitch         chan struct{}
	currConnection MyConn
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}

	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()

	<-s.quitch
	// close(s.MyConnCh)

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
		s.currConnection = *newConnection(conn, false)
		// conn.Write(MakeRespose("Hello Client!").GenerateBinaryMessage())
		go s.readLoop(s.currConnection.conn)
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

			s.currConnection.msgch <- Message{
				Sender: conn,
				Type:   string(command),
				Data:   string(message),
			}

			go s.handleMessage()

			foundLength = false
			messageLength = 0
		}
	}
}

func (s *Server) handleMessage() {
	for m := range s.currConnection.msgch {

		switch m.Type {
		case "LOGIN": // Login
			s.handleLogin(m)

		case "CREATEAUC": // Create Auction
			if !s.currConnection.loggedInStatus {
				m.Sender.Write(MakeRespose("You Must Login First").GenerateBinaryMessage())
			}

			s.handleCreateAuc(m)

		case "LISTAUC": // List Pending Auctions
			if !s.currConnection.loggedInStatus {
				m.Sender.Write(MakeRespose("You Must Login First").GenerateBinaryMessage())
			}

			s.handleListAuc(m)

		case "PLACEBID": // Place Bid
			if !s.currConnection.loggedInStatus {
				m.Sender.Write(MakeRespose("You Must Login First").GenerateBinaryMessage())
			}

			s.handlePlaceBid(m)

		case "ENDAUC": // End Auction
			if !s.currConnection.loggedInStatus {
				m.Sender.Write(MakeRespose("You Must Login First").GenerateBinaryMessage())
			}

			s.handleEndAuc(m)
		case "RETRIEVEBID": // Retrieve successful bids
			if !s.currConnection.loggedInStatus {
				m.Sender.Write(MakeRespose("You Must Login First").GenerateBinaryMessage())
			}

			s.handleRetrieveSuccBid(m)
		case "RETRIEVEWINNERS": // Retrieve winners of your auctions
			if !s.currConnection.loggedInStatus {
				m.Sender.Write(MakeRespose("You Must Login First").GenerateBinaryMessage())
			}

			s.handleRetrieveAucWinners(m)
		}
		// fmt.Printf("[%s]%s: %s\n", m.Sender.RemoteAddr(), m.Type, m.Data)
		// fmt.Print("Sending Confirmation ...\n\n")
		// m.Sender.Write(MakeRespose("Recieved your command " + m.Type).GenerateBinaryMessage())
	}
}
