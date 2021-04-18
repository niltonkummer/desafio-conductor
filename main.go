package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Bem-vindo %s", request.URL.Query().Get("name"))
	})

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadTimeout:       time.Second * 10,
		ReadHeaderTimeout: time.Second * 10,
		WriteTimeout:      time.Second * 10,
		IdleTimeout:       time.Second * 10,
		ErrorLog:          log.New(os.Stderr, "", log.LstdFlags),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
