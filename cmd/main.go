package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Doittikorn/cypernote/internal/infrastructure/config"
	"github.com/Doittikorn/cypernote/internal/infrastructure/server"
	"github.com/Doittikorn/cypernote/pkg/mlog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.New().All()
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Fatal(err.Error())
	}
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(mlog.Middleware(logger))
	e.Use(middleware.Recover())

	// Database
	sql, err := sql.Open("postgres", cfg.DBConnection)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// Ping database if err log and exit
	if err := sql.Ping(); err != nil {
		logger.Fatal(err.Error())
	}
	defer sql.Close()

	// Routes
	s := server.New(e, sql)
	s.RouteV1()

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil && err != http.ErrServerClosed {
			logger.Fatal("shutting down the server")
		}
		logger.Info("gracefully shutdown the server")
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	gCtx := context.Background()
	ctx, cancel := context.WithTimeout(gCtx, 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Fatal("unexpected shutdown the server", zap.Error(err))
	}
}
