// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/costaconrado/services-csm/config"
	v1 "github.com/costaconrado/services-csm/internal/controller/http/v1"
	"github.com/costaconrado/services-csm/internal/usecase/proposal"
	proposalRepo "github.com/costaconrado/services-csm/internal/usecase/proposal/repo"
	"github.com/costaconrado/services-csm/pkg/httpserver"
	"github.com/costaconrado/services-csm/pkg/logger"
	"github.com/costaconrado/services-csm/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	repo := proposalRepo.New(pg)
	proposalUseCase := proposal.New(repo)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, proposalUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
