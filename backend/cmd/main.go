package main

import (
	"coaching_backend/internal/config"
	"coaching_backend/internal/database"
	"coaching_backend/internal/server"
)

func main() {
	cnfg := config.LoadEnv()

	db := database.ConnectDB(cnfg)
	
	server.Start(db, cnfg)
}