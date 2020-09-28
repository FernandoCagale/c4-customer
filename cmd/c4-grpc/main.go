package main

import (
	"github.com/FernandoCagale/c4-customer/api/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	godotenv.Load()
}

func main() {
	session, err := SetupPostgres()
	if err != nil {
		panic("Erro to start Postgres")
	}

	defer session.Close()

	client, err := SetupClientGRPC()
	if err != nil {
		panic("Erro to start GRPC")
	}

	defer client.Close()

	events, err := SetupEvents(session, client)
	if err != nil {
		panic("Erro to start Events")
	}
	events.MakeEvents()

	app, err := SetupApplication(session, client)
	if err != nil {
		panic("Erro to start application")
	}

	router := app.MakeHandlers()

	router.Use(middleware.Header, middleware.Logging)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
