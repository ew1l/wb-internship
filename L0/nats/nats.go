package nats

import (
	"encoding/json"
	"log"
	"os"

	"L0/cache"
	"L0/database"
	"L0/models"

	"github.com/nats-io/stan.go"
)

var conn stan.Conn

func ConnectAndSubscribe() {
	var err error
	if conn, err = stan.Connect(
		os.Getenv("CLUSTER_ID"),
		os.Getenv("CLIENT_ID"),
		stan.NatsURL(stan.DefaultNatsURL),
		stan.SetConnectionLostHandler(func(_ stan.Conn, err error) {
			log.Println(err)
		}),
	); err != nil {
		log.Fatal(err)
	}

	if _, err = conn.Subscribe(os.Getenv("SUBJECT"), func(msg *stan.Msg) {
		log.Printf("Received on [%v]: %v", msg.Subject, string(msg.Data))
		if err := msg.Ack(); err != nil {
			log.Println(err)
		}

		order := models.Order{}
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			log.Println("Invalid data!")
		}

		if _, ok := cache.Get(order.OrderUID); !ok {
			database.Insert(order)
		}

	}, stan.SetManualAckMode()); err != nil {
		log.Println(err)
	}
}
