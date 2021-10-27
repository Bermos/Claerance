package main

import (
	"Claerance/internal/database"
	"Claerance/internal/entities/roles"
	"Claerance/internal/entities/sites"
	"Claerance/internal/entities/users"
	"Claerance/internal/server"
)

func main() {
	database.SetDriver("sqlite3")
	database.SetURI("test.db")

	users.Setup()
	sites.Setup()
	roles.Setup()

	server.InitSessionStore()
	server.Start(1401)
}
