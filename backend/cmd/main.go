package main

import (
	"coaching_backend/internal/config"
	"coaching_backend/internal/database"
	"coaching_backend/internal/server"
	"fmt"
)

func main() {
	cnfg, err := config.LoadEnv()

	if err != nil {
		fmt.Println("env. load error :",err.Error())
	}

	db := database.ConnectDB(cnfg)
	
	server.Start(db, cnfg)
}