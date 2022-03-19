package cache

import "github.com/ew1l/wb-l0/models"

var cache map[string]models.Order

func Init() {
	cache = make(map[string]models.Order, 0)
}

func Insert(order models.Order) {
	cache[order.OrderUID] = order
}

func Get(id string) (models.Order, bool) {
	order, ok := cache[id]

	return order, ok
}
