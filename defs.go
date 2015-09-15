package owfs

import (
	"log"
)

const (
	msg_error       = iota
	msg_nop         = iota
	msg_read        = iota
	msg_write       = iota
	msg_dir         = iota
	msg_size        = iota // unused
	msg_presence    = iota
	msg_dirall      = iota
	msg_get         = iota
	msg_dirallslash = iota
	msg_getslash    = iota
)

type RequestHeader struct {
	Version       uint32
	PayloadLength uint32
	Type          uint32
	Flags         uint32
	Size          uint32
	Offset        uint32
}
type ResponseHeader struct {
	Version       uint32
	PayloadLength uint32
	Ret           uint32
	Flags         uint32
	Size          uint32
	Offset        uint32
}

func (h ResponseHeader) dump() {
	log.Println("Response Header")
	log.Println("   Version:       ", h.Version)
	log.Println("   PayloadLength: ", h.PayloadLength)
	log.Println("   Ret:           ", h.Ret)
	log.Println("   Flags:         ", h.Flags)
	log.Println("   Size:          ", h.Size)
	log.Println("   Offset:        ", h.Offset)
}
