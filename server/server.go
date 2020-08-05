package server

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	logrusMiddleware "github.com/neko-neko/echo-logrus/v2"
	logrusLogger "github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"
	"go-clean-architecture/internal/controller"
	"os"
	"os/signal"
	"time"
)

// Configs defines the application configuration
type Configs struct {
	Port       string      `json:"port"`
	LogLVL     string      `json:"log"`
}

// Start starts the server
func Start(c Configs) error {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Logger = logrusLogger.Logger()
	e.Use(logrusMiddleware.Logger(),
		middleware.Recover())

	lvl, err := logrus.ParseLevel(c.LogLVL)
	if err == nil {
		logrus.SetLevel(lvl)
	}

	api := controller.NewAPI()
	api.RegisterHandlers(e)


	// Let users know the running configuration
	configJSON, err := json.Marshal(c)
	if err != nil {
		e.Logger.Error("Failed to print configuration", err)
	}
	e.Logger.Info("Configuration: ", string(configJSON))

	// Start server
	go func() {
		if err := e.Start(":" + c.Port); err != nil {
			e.Logger.Error("shutting down the server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
