package server

import (
	"log"
	"net/http"
	"text/template"

	"L0/cache"
	"L0/database"
)

func Run() {
	defer database.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/order", OrderHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func OrderHandler(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	order, ok := cache.Get(id)
	if !ok {
		http.NotFound(rw, r)
		return
	}

	template, err := template.ParseFiles("web/templates/order.html")
	if err != nil {
		log.Println(err)
		return
	}

	if err = template.Execute(rw, order); err != nil {
		log.Println(err)
	}
}
