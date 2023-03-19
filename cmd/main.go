package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Doittikorn/cypernote/internal/app"
	"github.com/Doittikorn/cypernote/internal/config"
	"github.com/Doittikorn/cypernote/pkg/mlog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

var (
	// database connection
	database *sql.DB
)

func main() {
	// Initialize Config
	cfg := config.New().All()
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Initialize Database Connection
	database, err := sql.Open("postgres", cfg.DBConnection)
	if err != nil {
		panic(err)
	}
	if err := database.Ping(); err != nil {
		logger.Fatal(err.Error())
	}
	defer database.Close()

	// API
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(mlog.Middleware(logger))
	e.Use(middleware.Recover())

	// Routes
	// s := route.New(e, sql)
	// s.RouteV1()
	// s.Health()

	// Instance App
	app := app.InitializeApp(database)

	// Routes
	e.POST("v1/finance/", app.Handler.Finance.Save)
	e.DELETE("v1/finance/", app.Handler.Finance.Delete)
	e.GET("v1/finance/user/:id", app.Handler.Finance.GetByUserID)

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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logger.Fatal("unexpected shutdown the server", zap.Error(err))
	}
}
