package main

import (
	"Claerance/internal/database"
	"Claerance/internal/server"
	"Claerance/internal/users"
)

func main() {
	database.SetDriver("sqlite3")
	database.SetURI("test.db")

	users.Setup()

	server.InitSessionStore()
	server.Start(1401)
}
