package server

import (
	"fmt"
	"log"
	"net"

	"github.com/Sp33ktrE/redis-clone/resp"
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
		resp := resp.New(conn)

		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(value)

		conn.Write([]byte("+OK\r\n"))
	}
}
