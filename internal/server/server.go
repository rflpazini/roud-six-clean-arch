package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"rflpazini/round-six/config"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

type Server struct {
	echo   *echo.Echo
	config *config.Config
	logger *zap.Logger
}

func NewServer(cfg *config.Config, log *zap.Logger) *Server {
	return &Server{
		echo.New(),
		cfg,
		log,
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
		s.logger.Sugar().Infof("Server is listening on PORT: %s", s.config.Server.Port)
		if err := s.echo.StartServer(server); err != nil {
			s.logger.Sugar().Fatalf("Error starting Server: ", err)
		}
	}()

	go func() {
		s.logger.Sugar().Infof("Starting Debug Server on PORT: %s", s.config.Server.PprofPort)
		if err := http.ListenAndServe(s.config.Server.PprofPort, http.DefaultServeMux); err != nil {
			s.logger.Sugar().Errorf("Error PPROF ListenAndServe: %s", err)
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

	s.logger.Sugar().Info("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}
