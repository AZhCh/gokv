package core

import (
	"bufio"
	"io"
	"log"
	"net"
)

type MainHandler interface {
	Handle(conn net.Conn)
}

func NewMainHandler() MainHandler {
	return &connHandler{}
}

type connHandler struct {
}

func (c connHandler) Handle(conn net.Conn) {
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	bs := make([]byte, 1024)

	reader := bufio.NewReader(conn)
	for {
		_, err := reader.Read(bs)

		if err == io.EOF {
			log.Printf("client closed connection: %s", conn.RemoteAddr().String())
			break
		}
	}
}
