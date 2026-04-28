package httpserver

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/re-partners-challenge-backend/internal/infra/config"
	"go.uber.org/zap"
)

type Server struct {
	config           *config.ServerConfig
	handler          http.Handler
	logger           *zap.Logger
	server           *http.Server
	onBootstrapHooks []func(ctx context.Context) error
	onShutDownHooks  []func(ctx context.Context) error
}

type ServerOption struct {
	Config  *config.Config
	Handler http.Handler
	Logger  *zap.Logger
}

func ProvideHTTPServer(opt *ServerOption) *Server {
	return &Server{
		config:  &opt.Config.Server,
		handler: opt.Handler,
		logger:  opt.Logger,
		server: &http.Server{
			Addr:         opt.Config.Server.Address(),
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      opt.Handler,
		},
		onBootstrapHooks: make([]func(ctx context.Context) error, 0),
		onShutDownHooks:  make([]func(ctx context.Context) error, 0),
	}
}

func (s *Server) OnBootstrap(f ...func(ctx context.Context) error) *Server {
	s.onBootstrapHooks = append(s.onBootstrapHooks, f...)
	return s
}

func (s *Server) OnShutDown(f ...func(ctx context.Context) error) *Server {
	s.onShutDownHooks = append(s.onShutDownHooks, f...)
	return s
}

func (s *Server) Start() error {

	backgroundCtx := context.Background()
	for _, onBootstrap := range s.onBootstrapHooks {
		if err := onBootstrap(backgroundCtx); err != nil {
			s.logger.Error("a bootstrap func returned an error", zap.Error(err))
		}
	}

	s.logger.Info("Server is running", zap.String("address", s.config.Address()))

	errorChannel := make(chan error, 1)
	stopSignalChannel := make(chan os.Signal, 1)

	signal.Notify(stopSignalChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGABRT)

	go func() {
		errorChannel <- s.server.ListenAndServe()
	}()

	select {
	case err := <-errorChannel:

		s.logger.Error("failed to start the server listener", zap.Error(err))

		return err

	case sign := <-stopSignalChannel:

		s.logger.Info("received signal to stop the server", zap.String("signal", sign.String()))

		ctx, cancel := context.WithTimeout(backgroundCtx, time.Minute*5)
		defer cancel()

		if err := s.server.Shutdown(ctx); err != nil {
			s.logger.Error("failed to stop the server listener", zap.Error(err))
			return err
		}

		for _, onShutDown := range s.onShutDownHooks {
			if err := onShutDown(ctx); err != nil {
				s.logger.Error("a shutdown func returned a error", zap.Error(err))
			}
		}

		s.logger.Info("http server stopped successfully")

		return nil
	}

}
