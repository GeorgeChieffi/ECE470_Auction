package message

import "encoding/binary"

type Message struct {
	Type      string
	Data      string
	binaryMsg []byte
}

func CreateMessage(sender string, reciever string, cmd string, Data string) *Message {
	return &Message{
		Type:      cmd,
		Data:      Data,
		binaryMsg: GenerateBinaryMessage(cmd, Data),
	}
}

func DecodeMessage([]byte) *Message {

	return &Message{
		Sender:
	}
}

func GenerateBinaryMessage(Type string, Data string) []byte {
	maxDataLen := 4
	maxCMDLen := 2
	dataBuff := make([]byte, maxDataLen)
	CMDBuff := make([]byte, maxCMDLen)

	binary.LittleEndian.PutUint32(dataBuff, uint32(len(Data)))
	binary.LittleEndian.PutUint16(CMDBuff, uint16(len(Type)))

	bytes := append(append(append(dataBuff, CMDBuff...), []byte(Type)...), []byte(Data)...)

	return bytes
}

