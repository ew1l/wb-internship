package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/ew1l/wb-l0/cache"
	"github.com/ew1l/wb-l0/models"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func Close() {
	db.Close()
}

func Insert(order models.Order) {
	query_orders := `INSERT INTO orders(uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	query_delivery := `INSERT INTO delivery(name, phone, zip, city, address, region, email, order_uid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	query_payments := `INSERT INTO payments(transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, order_uid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	query_items := `INSERT INTO items(chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, order_uid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	db.QueryRow(query_orders, order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.Shardkey, order.SmID, order.DateCreated, order.OofShard)
	db.QueryRow(query_delivery, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email, order.OrderUID)
	db.QueryRow(query_payments, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee, order.OrderUID)

	for _, item := range order.Items {
		db.QueryRow(query_items, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status, order.OrderUID)
	}

	cache.Insert(order)
}

func RestoreCache() {
	cache.Init()

	query_orders := `SELECT * FROM orders`
	query_delivery := `SELECT name, phone, zip, city, address, region, email FROM delivery WHERE order_uid = $1`
	query_payments := `SELECT transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee FROM payments WHERE order_uid = $1`
	query_items := `SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status FROM items WHERE order_uid = $1`

	rows_orders, err := db.Query(query_orders)
	if err != nil {
		log.Println(err)
	}
	defer rows_orders.Close()

	for rows_orders.Next() {
		order := models.Order{}
		rows_orders.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature, &order.CustomerID, &order.DeliveryService, &order.Shardkey, &order.SmID, &order.DateCreated, &order.OofShard)

		rows_delivery, err := db.Query(query_delivery, order.OrderUID)
		if err != nil {
			log.Println(err)
		}
		defer rows_delivery.Close()

		delivery := models.Delivery{}
		for rows_delivery.Next() {
			rows_delivery.Scan(&delivery.Name, &delivery.Phone, &delivery.Zip, &delivery.City, &delivery.Address, &delivery.Region, &delivery.Email)
		}
		order.Delivery = delivery

		rows_payments, err := db.Query(query_payments, order.OrderUID)
		if err != nil {
			log.Println(err)
		}
		defer rows_payments.Close()

		payment := models.Payment{}
		for rows_payments.Next() {
			rows_payments.Scan(&payment.Transaction, &payment.RequestID, &payment.Currency, &payment.Provider, &payment.Amount, &payment.PaymentDT, &payment.Bank, &payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee)
		}
		order.Payment = payment

		rows_items, err := db.Query(query_items, order.OrderUID)
		if err != nil {
			log.Println(err)
		}
		defer rows_items.Close()

		items := make([]models.Item, 0)
		for rows_items.Next() {
			item := models.Item{}
			rows_items.Scan(&item.ChrtID, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status)

			items = append(items, item)
		}
		order.Items = items

		cache.Insert(order)
	}
}
