package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/alesabater/idea-tracker/idea-api/handlers"
)

func main() {
	fmt.Println("Started ideas server")

	l := log.New(os.Stdout, "idea-api", log.LstdFlags)
	ig := handlers.NewIdeaService(l)

	sm := http.NewServeMux()
	sm.Handle("/", ig)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      ig,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt) // broadcast event on c channel whenever Interrupt happens
	signal.Notify(c, os.Kill)      // broadcast event on c channel whenever Kill happens

	sig := <-c
	l.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}
