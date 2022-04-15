package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rss_reader_service/internal/http/rest"
	"rss_reader_service/internal/service"
	"syscall"
	"time"
)

func main() {

	readerService := service.NewRssReaderService()
	handler := rest.NewHandler(readerService)

	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/parse", handler.ParseURLs)

	errs := make(chan error, 2)
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "9000"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {

		fmt.Println("Listening on port : " + port)
		errs <- server.ListenAndServe()

	}()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer func() {
		cancel()
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT)

	errs <- fmt.Errorf("%s", <-c)

	fmt.Printf("Graceful shutdown %s", <-errs)

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

}
