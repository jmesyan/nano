package game

import (
	"bytes"
	"encoding/binary"
)

// Packet represents a network packet.
type Packet struct {
	Length int32
	Data   []byte
	Cid    int32
	Cmd    int32
	N      int32
	T      int32
}

func NewPacket(data []byte, size int32) *Packet {
	var cid, cmd, nn, tt int32
	bbuf := bytes.NewBuffer(data[0:4])
	binary.Read(bbuf, binary.LittleEndian, &cid)
	bbuf = bytes.NewBuffer(data[4:8])
	binary.Read(bbuf, binary.LittleEndian, &cmd)
	bbuf = bytes.NewBuffer(data[8:12])
	binary.Read(bbuf, binary.LittleEndian, &nn)
	bbuf = bytes.NewBuffer(data[12:16])
	binary.Read(bbuf, binary.LittleEndian, &tt)
	udata := data[16:size]
	return &Packet{
		Length: size - 16,
		Data:   udata,
		Cid:    cid,
		Cmd:    cmd,
		N:      nn,
		T:      tt,
	}
}
