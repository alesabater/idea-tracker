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
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	gohandlers "github.com/gorilla/handlers"
)

func main() {
	fmt.Println("Started ideas server")

	l := log.New(os.Stdout, "idea-api", log.LstdFlags)
	ig := handlers.NewIdeaService(l)

	// root router
	sm := mux.NewRouter()

	// turns subrouter into a router
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ig.GetIdeas)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ig.UpdateIdea)
	putRouter.Use(ig.MiddlewareIdeaValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ig.AddIdea)
	postRouter.Use(ig.MiddlewareIdeaValidation)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", ig.DeleteIdea)
	//sm.Handle("/ideas", ig)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	sm = mux.NewRouter()

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	s := &http.Server{
		Addr:         ":9091",
		Handler:      ch(sm),
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
