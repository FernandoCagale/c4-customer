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

	events, err := SetupEvents(session)
	if err != nil {
		panic("Erro to start Events")
	}
	events.MakeEvents()

	app, err := SetupApplication(session)
	if err != nil {
		panic("Erro to start application")
	}

	router := app.MakeHandlers()

	router.Use(middleware.Header)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
