package main

import "example/event-board/server"

func main() {
	server := server.InitServer()
	server.Run()
}
