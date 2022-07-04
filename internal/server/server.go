package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"rflpazini/round-six/config"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

type Server struct {
	echo   *echo.Echo
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		echo.New(),
		cfg,
	}
}

func (s *Server) Run() error {

	server := &http.Server{
		Addr:           s.config.Server.Port,
		ReadTimeout:    time.Second * s.config.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.config.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		log.Infof("Server is listening on PORT: %s", s.config.Server.Port)
		if err := s.echo.StartServer(server); err != nil {
			log.Fatalf("Error starting Server: ", err)
		}
	}()

	go func() {
		log.Infof("Starting Debug Server on PORT: %s", s.config.Server.PprofPort)
		if err := http.ListenAndServe(s.config.Server.PprofPort, http.DefaultServeMux); err != nil {
			log.Errorf("Error PPROF ListenAndServe: %s", err)
		}
	}()

	if err := s.Handlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	log.Info("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}
