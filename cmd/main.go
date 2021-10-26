package main

import (
	"Claerance/internal/database"
	"Claerance/internal/roles"
	"Claerance/internal/server"
	"Claerance/internal/sites"
	"Claerance/internal/users"
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
