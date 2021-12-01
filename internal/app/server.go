package app

import (
	"location-history-server/internal/data"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Config holds the attributes of the Server
type Config struct {
	Name        string
	Port        string
	UseInMemory bool
}

func newConfig() *Config {
	c := new(Config)

	c.Name = "location-history"
	c.Port = mustReadPort()
	c.UseInMemory = mustReadInMemoryMode()
	return c
}

// Server is the top level location-history server application object.
type Server struct {
	// TODO : db repo
	*Config

	repo data.Repo
}

// NewServer creates server object
func NewServer() *Server {
	s := new(Server)

	s.Config = newConfig()
	s.repo = data.NewRepo(s.UseInMemory)
	return s
}

func (s *Server) Start() {
	server := http.Server{
		Addr:    s.Port,
		Handler: s.InitRouter(),
	}

	go func() {
		log.Printf("listening on port %s...\n", server.Addr)
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal("failed to start server: ", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}
