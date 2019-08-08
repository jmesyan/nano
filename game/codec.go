package game

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// Codec constants.
const (
	HeadLength    = 4
	MaxPacketSize = 64 * 1024
)

// ErrPacketSizeExcced is the error used for encode/decode.
var ErrPacketSizeExcced = errors.New("codec: packet size exceed")

// A Decoder reads and decodes network data slice
type Decoder struct {
	buf  *bytes.Buffer
	size int32 // last packet length
}

// NewDecoder returns a new decoder that used for decode network bytes slice.
func NewDecoder() *Decoder {
	return &Decoder{
		buf:  bytes.NewBuffer(nil),
		size: -1,
	}
}

func (c *Decoder) forward() error {
	header := c.buf.Next(HeadLength)
	bbuf := bytes.NewBuffer(header)
	binary.Read(bbuf, binary.LittleEndian, &c.size)
	c.size += 16
	// packet length limitation
	if c.size > MaxPacketSize {
		return ErrPacketSizeExcced
	}
	return nil
}

// Decode decode the network bytes slice to Packet(s)
// TODO(Warning): shared slice
func (c *Decoder) Decode(data []byte) ([]*Packet, error) {
	c.buf.Write(data)

	var (
		packets []*Packet
		err     error
	)
	// check length
	if c.buf.Len() < HeadLength {
		return nil, err
	}

	// first time
	if c.size < 0 {
		if err = c.forward(); err != nil {
			return nil, err
		}
	}

	for c.size <= int32(c.buf.Len()) {
		p := NewPacket(c.buf.Next(int(c.size)), c.size)
		packets = append(packets, p)

		// more packet
		if c.buf.Len() < HeadLength {
			c.size = -1
			break
		}

		if err = c.forward(); err != nil {
			return nil, err
		}
	}

	return packets, nil
}
