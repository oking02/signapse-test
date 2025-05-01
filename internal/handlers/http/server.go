package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	port   int
	routes map[string]http.Handler
}

func NewServer(port int) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) SetupRoutes(routes map[string]http.Handler) {
	s.routes = routes
}

func (s *Server) Run(ctx context.Context) error {

	ctx, cancel := context.WithCancel(ctx)

	mux := http.NewServeMux()
	mux.HandleFunc("/", base)

	if s.routes != nil {
		for path, handler := range s.routes {
			mux.Handle(path, handler)
		}
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: mux,
	}

	go func() {
		defer cancel()
		log.Printf("starting server on :%d", s.port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	ctx, shutdownCancel := context.WithTimeout(ctx, 5*time.Second)
	defer shutdownCancel()
	if err := server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

func base(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!"))
}
