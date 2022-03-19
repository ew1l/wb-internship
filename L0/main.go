package main

import (
	"log"

	"L0/database"
	"L0/nats"
	"L0/web/server"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	database.Connect()
	database.RestoreCache()

	nats.ConnectAndSubscribe()
	server.Run()
}
