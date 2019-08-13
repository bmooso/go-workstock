package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bmooso/go-workstock/handlers"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.HideBanner = true

	h := handlers.Handler{}

	h.InitRoutes(e)

	// ------------------------------------------------------
	// server
	// ------------------------------------------------------
	go func() {
		if err := e.Start(":8080"); err != nil {
			fmt.Println("shutting down the server")
		}
	}()

	// ------------------------------------------------------
	// shutdown
	// ------------------------------------------------------
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
