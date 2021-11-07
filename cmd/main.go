package main

import (
	"Claerance/internal/config"
	"Claerance/internal/database"
	"Claerance/internal/entities/roles"
	"Claerance/internal/entities/sites"
	"Claerance/internal/entities/users"
	"Claerance/internal/server"
)

func main() {
	config.Load()

	database.SetDriver("sqlite3")
	database.SetURI("test.db")

	users.Setup()
	sites.Setup()
	roles.Setup()

	server.InitSessionStore()
	server.Start()
}
