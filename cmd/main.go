package main

import (
	"Claerance/internal/database"
	"Claerance/internal/server"
	userManger "Claerance/internal/users/manager"
)

func main() {
	database.SetDriver("sqlite3")
	database.SetURI("test.db")

	userManger.Setup()

	server.InitSessionStore()
	server.Start(1401)
}
