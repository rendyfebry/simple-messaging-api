package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	httptransport "github.com/rendyfebry/simple-messaging-api/transport/http"
)

func main() {
	routes := httptransport.MakeRoutes()

	// Routes here

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%v", "0.0.0.0", "8080"),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      routes,
	}

	fmt.Println(fmt.Sprintf("Environment: %s", "development"))
	fmt.Println(fmt.Sprintf("Application URL: http://%s:%v", "localhost", "8080"))

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
