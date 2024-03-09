package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (s *Server) handleLogin(m Message) {
	var data = m.Data

	// fmt.Println(data)

	// Split the string by comma delimiter
	substrings := strings.Split(data, "&")

	u := substrings[0]
	p := substrings[1]

	var err error
	var res string
	s.currConnection.user, err = getUser(u, p)
	if err != nil {
		res = "LOGIN_DENIED"
	} else {
		res = "LOGIN_SUCCESS"
		s.currConnection.loggedInStatus = true
	}

	m.Sender.Write(MakeRespose(res).GenerateBinaryMessage())
}

func (s *Server) handleCreateAuc(m Message) {
	var res string
	strs := strings.Split(m.Data, "&")
	if len(strs) != 7 {
		res = "CREATEAUC_DENIED"
	}
	p, _ := strconv.ParseFloat(strs[6], 64)
	insertAuc(strs[0], strs[1], strs[2], strs[3], strs[4], strs[5], p, s.currConnection.user.UID)
	res = "CREATEAUC_SUCCESSFUL"

	m.Sender.Write(MakeRespose(res).GenerateBinaryMessage())
}

func (s *Server) handleListAuc(m Message) {
	var res string
	var err error

	res, err = getAucs()
	if err != nil {
		res = fmt.Sprintf("ERROR: %s", err)
	}

	m.Sender.Write(MakeRespose(res).GenerateBinaryMessage())
}

func (s *Server) handlePlaceBid(m Message) {
	var res string
	substrings := strings.Split(m.Data, "&")
	if len(substrings) != 2 {
		res = "PLACEBID_DENIED"
	}
	createBid(substrings[0], substrings[1], s.currConnection.user.UID)
	res = "PLACEBID_SUCCESSFUL"
	m.Sender.Write(MakeRespose(res).GenerateBinaryMessage())
}

func (s *Server) handleEndAuc(m Message) {
	var res string
	err := endAucDB(s.currConnection.user.UID, m.Data)
	if err != nil {
		res = "ENDAUC_DENIED"
	}
	res = "ENDAUC_SUCCESSFUL"

	m.Sender.Write(MakeRespose(res).GenerateBinaryMessage())
}

func (s *Server) handleRetrieveSuccBid(m Message) {
	var res string

	res, err := getWinningBids(s.currConnection.user.UID)
	if err != nil {
		res = fmt.Sprintf("ERROR: %s", err)
	}

	m.Sender.Write(MakeRespose(res).GenerateBinaryMessage())
}

func (s *Server) handleRetrieveAucWinners(m Message) {
	var res string

	res, err := getWinners(s.currConnection.user.UID)
	if err != nil {
		res = fmt.Sprintf("ERROR: %s", err)
	}

	m.Sender.Write(MakeRespose(res).GenerateBinaryMessage())
}
