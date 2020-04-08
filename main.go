package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rendyfebry/simple-messaging-api/services"
	httptransport "github.com/rendyfebry/simple-messaging-api/transport/http"
)

func main() {
	// Initialize service
	svc := services.NewService("local")

	// Initialize route
	routes := httptransport.MakeRoutes(svc)

	// Should coming from env or config file
	host := "0.0.0.0"
	port := 8080

	// Initialize http serve
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      routes,
	}

	fmt.Println(fmt.Sprintf("Environment: %s", "development"))
	fmt.Println(fmt.Sprintf("Application URL: http://%s:%d", host, port))

	// Run server in goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
