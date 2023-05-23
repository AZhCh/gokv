package server

import (
	"gokv/core"
	"log"
	"net"
)

type Server interface {
	Start()
}

type TcpServer struct {
	Port int
}

func (tl *TcpServer) Start() {
	tcpAddr := net.TCPAddr{
		IP:   nil,
		Port: tl.Port,
		Zone: "",
	}
	tcp, err := net.ListenTCP("tcp", &tcpAddr)
	defer func(tcp *net.TCPListener) {
		err := tcp.Close()
		if err != nil {
			log.Fatalf("quit error: %v", err)
		}
	}(tcp)
	if err != nil {
		panic(err)
	}

	handler := core.NewMainHandler()
	for {
		conn, err := tcp.AcceptTCP()
		if err != nil {
			log.Printf("accept tcp: %v", err)
			continue
		}

		go handler.Handle(conn)
	}
}
