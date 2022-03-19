package main

import (
	"log"

	"github.com/ew1l/wb-l0/database"
	"github.com/ew1l/wb-l0/nats"
	"github.com/ew1l/wb-l0/web/server"
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
