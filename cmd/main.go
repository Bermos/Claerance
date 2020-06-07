package main

import (
	"Claerance/internal/server"
	"Claerance/internal/users"
)

func main() {
	users.Setup()
	server.Start(1401)
}
