package main

import (
	"encoding/binary"
)

type Message struct {
	Type string
	Data string
}

func MakeMessage(cmd string, Data string) *Message {
	return &Message{
		Type: cmd,
		Data: Data,
	}
}

func (m *Message) GenerateBinaryMessage() []byte {
	maxDataLen := 4
	maxCMDLen := 2
	dataBuff := make([]byte, maxDataLen)
	CMDBuff := make([]byte, maxCMDLen)

	binary.LittleEndian.PutUint32(dataBuff, uint32(len(m.Data)))
	binary.LittleEndian.PutUint16(CMDBuff, uint16(len(m.Type)))

	bytes := append(append(append(dataBuff, CMDBuff...), []byte(m.Type)...), []byte(m.Data)...)

	return bytes
}
