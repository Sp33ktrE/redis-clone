package main

import (
	"github.com/Sp33ktrE/redis-clone/server"
)

const HOST = ""
const PORT = "6379"

func main() {
	server := server.New(HOST, PORT)
	server.Run()
}
