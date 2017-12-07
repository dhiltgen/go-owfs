package owfs

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

// Read raw byte array from a 1-wire device
func (oc *OwfsClient) Read(path string) (string, error) {
	conn, err := net.Dial("tcp", oc.connString)
	if err != nil {
		return "", fmt.Errorf("Failed to connect to owserver: %s", err)
	}
	defer conn.Close()

	payload := make([]byte, len(path)+1) // To get null terminated
	copy(payload, []byte(path))
	data := RequestHeader{
		Type:          msg_read,
		PayloadLength: uint32(len(payload)),
		Size:          65536,
	}

	//log.Println("Attempting to send data")
	err = binary.Write(conn, binary.BigEndian, data)
	if err != nil {
		return "", fmt.Errorf("Failed to write header to owserver: %s", err)
	}
	err = binary.Write(conn, binary.BigEndian, payload)
	if err != nil {
		return "", fmt.Errorf("Failed to write payload to owserver: %s", err)
	}

	//log.Println("Attempting to read response")
	var response ResponseHeader

	err = binary.Read(conn, binary.BigEndian, &response)
	if err != nil {
		return "", fmt.Errorf("Failed to read header from owserver: %s", err)
	}

	if response.PayloadLength > 0 {
		buf := make([]byte, response.PayloadLength)
		err = binary.Read(conn, binary.BigEndian, &buf)
		if err != nil {
			return "", fmt.Errorf("Failed to read payload from owserver: %s", err)
		}
		return strings.TrimSpace(string(buf)), nil
	} else {
		//response.dump()
		return "", fmt.Errorf("Zero length data for device %s", path)
	}
}
