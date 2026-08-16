package main

import (
	"coaching_backend/internal/database"
	"coaching_backend/internal/server"
)

func main() {
	

	server.Start()
	database.ConnectDB()

}