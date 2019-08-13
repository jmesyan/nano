package apiserver

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmesyan/nano"
	"github.com/jmesyan/nano/utils"
)

// Codec constants.
const (
	HeadLength    = 32
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
func (c *Decoder) Decode(data []byte) (*Packet, error) {
	var (
		packet *Packet
		err    error
	)
	strData := string(data)
	// check length
	lenData := len(strData)
	if lenData < HeadLength {
		return nil, errors.New("the str data is too short")
	}
	strSign := strData[0:32]
	strJson := strData[32:]
	nano.logger.Println(">> api:", strJson)
	sign := utils.Md5(fmt.Sprintf("go-realbull-api%s", strJson))
	if sign != strSign {
		return nil, errors.New("the sign is error")
	}
	obj := make(map[string]interface{})
	err = json.Unmarshal([]byte(strJson), obj)
	if err != nil {
		return nil, err
	}
	packet = NewPacket(obj, lenData)
	return packet, nil
}
