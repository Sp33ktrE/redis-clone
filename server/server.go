package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type Server struct {
	host string
	port string
}

func New(host string, port string) *Server {
	return &Server{
		host: host,
		port: port,
	}
}

func (server *Server) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// read client message
	for {
		buff := make([]byte, 1024)

		_, err := conn.Read(buff)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading from client: ", err.Error())
			os.Exit(1)
		}

		conn.Write([]byte("+OK\r\n"))
	}
}
