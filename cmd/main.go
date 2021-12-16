package main

import (
	"Claerance/internal/config"
	"Claerance/internal/database"
	"Claerance/internal/schemas"
	"Claerance/internal/server"
)

func main() {
	config.Load()

	database.SetDriver("sqlite3")
	database.SetURI("test.db")

	schemas.Setup()

	server.InitSessionStore()
	server.Start()
}
