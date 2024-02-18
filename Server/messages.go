package main

import (
	"encoding/binary"
)

type Response struct {
	Data string
}

func MakeRespose(Data string) *Response {
	return &Response{
		Data: Data,
	}
}

func (m *Response) GenerateBinaryMessage() []byte {
	maxDataLen := 4
	dataBuff := make([]byte, maxDataLen)

	binary.LittleEndian.PutUint32(dataBuff, uint32(len(m.Data)))

	bytes := append(dataBuff, []byte(m.Data)...)

	return bytes
}
