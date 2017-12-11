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
	Version       int32
	PayloadLength int32
	Type          int32
	Flags         int32
	Size          int32
	Offset        int32
}
type ResponseHeader struct {
	Version       int32
	PayloadLength int32
	Ret           int32
	Flags         int32
	Size          int32
	Offset        int32
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
