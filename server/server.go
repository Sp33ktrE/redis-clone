package server

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/Sp33ktrE/redis-clone/cmd"
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
		respReader := resp.NewReader(conn)

		value, err := respReader.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		if value.Typ != "array" {
			fmt.Println("Invalid request, expected array")
			continue
		}

		if len(value.Array) == 0 {
			fmt.Println("Invalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]

		respWriter := resp.NewWriter(conn)

		handler, ok := cmd.Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			respWriter.Write(resp.Value{Typ: "string", Str: ""})
			continue
		}
		result := handler(args)
		respWriter.Write(result)
	}
}
