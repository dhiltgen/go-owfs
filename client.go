package owfs

import (
	"log"
	"net"
)

type OwfsClient struct {
	connString string
	conn       net.Conn // TODO - Or should we re-establish on every request?
}

func NewClient(connString string) (OwfsClient, error) {
	conn, err := net.Dial("tcp", connString)
	if err != nil {
		log.Fatalf("Failed to connect to owserver: %s", err)
	}
	oc := OwfsClient{
		connString: connString,
		conn:       conn,
	}
	return oc, nil
}

func (oc *OwfsClient) Close() {
	oc.conn.Close()
}
