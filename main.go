package main

import (
	"context"
	"golang-microservice/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main(){
	lgr := log.New(os.Stdout, "product-api",log.LstdFlags)
	ph := handlers.NewProducts(lgr)
	sm := http.NewServeMux()
	sm.Handle("/", ph)
	s := &http.Server{
		Addr: ":8000",
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler: sm,
	}
	go func(){
		lgr.Println("Starting server on 8000 port!!")
		err := s.ListenAndServe()
		if err != nil {
			lgr.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	sig := <-c
	lgr.Println("Got signal: ", sig)
	tc, cancel := context.WithDeadline(context.Background(),  time.Now().Add(30 * time.Second))
	defer cancel()
	s.Shutdown(tc)
}