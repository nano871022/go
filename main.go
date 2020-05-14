package main

import (
	"./servers"
)

var port = 9192

func main() {
	servers.StartServer(port)
}
