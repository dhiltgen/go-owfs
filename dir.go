package owfs

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

func (oc *OwfsClient) Dir(path string) ([]string, error) {
	conn, err := net.Dial("tcp", oc.connString)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to owserver: %s", err)
	}
	defer conn.Close()

	payload := make([]byte, len(path)+1) // To get null terminated
	copy(payload, []byte(path))
	data := RequestHeader{
		Type:          msg_dir,
		PayloadLength: uint32(len(payload)),
	}

	ret := []string{}

	//log.Println("Attempting to send data")
	err = binary.Write(conn, binary.BigEndian, data)
	if err != nil {
		return nil, fmt.Errorf("Failed to write header to owserver: %s", err)
	}
	err = binary.Write(conn, binary.BigEndian, payload)
	if err != nil {
		return nil, fmt.Errorf("Failed to write payload to owserver: %s", err)
	}

	//log.Println("Attempting to read response")
	var response ResponseHeader

	for {
		err = binary.Read(conn, binary.BigEndian, &response)
		if err != nil {
			return nil, fmt.Errorf("Failed to read header from owserver: %s", err)
		}

		//response.dump()
		if response.PayloadLength > 0 {
			buf := make([]byte, response.PayloadLength)
			err = binary.Read(conn, binary.BigEndian, &buf)
			if err != nil {
				return nil, fmt.Errorf("Failed to read payload from owserver: %s", err)
			}
			ret = append(ret, strings.TrimSpace(string(buf[:len(buf)-1])))
			//log.Println("Response Payload: ", string(buf))
		} else {
			break
		}
	}
	return ret, nil
}
