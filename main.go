package main

import "gokv/server"

func main() {
	s := server.TcpServer{
		Port: 8080,
	}

	s.Start()
}
