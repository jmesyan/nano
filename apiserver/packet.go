package apiserver

// Packet represents a network packet.
type Packet struct {
	Length int
	Data   map[string]interface{}
}

func NewPacket(data map[string]interface{}, size int) *Packet {
	return &Packet{
		Length: size,
		Data:   data,
	}
}
