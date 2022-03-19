package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"L2/develop/dev11/handler"
	"L2/develop/dev11/service"

	"github.com/gorilla/mux"
)

// Server structure
type Server struct {
	*Config
}

// NewServer server's structure constructor
func NewServer() *Server {
	return &Server{
		Config: NewConfig(),
	}
}

// Run starts the server
func (s *Server) Run() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT)

	router := mux.NewRouter()
	service := service.NewService()
	handler := handler.NewHandler(service)

	router.HandleFunc("/create_event", handler.CreateEvent).Methods("POST")
	router.HandleFunc("/update_event", handler.UpdateEvent).Methods("POST")
	router.HandleFunc("/delete_event", handler.DeleteEvent).Methods("POST")
	router.HandleFunc("/events_for_day", handler.EventsForDay).Methods("GET")
	router.HandleFunc("/events_for_week", handler.EventsForWeek).Methods("GET")
	router.HandleFunc("/events_for_month", handler.EventsForMonth).Methods("GET")

	server := &http.Server{
		Addr:    net.JoinHostPort(s.Config.Host, s.Config.Port),
		Handler: WithLogging(router),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}
