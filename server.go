package main

import (
	"context"
	"film36exp/db"
	"film36exp/routes"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	negronilogrus "github.com/meatballhat/negroni-logrus"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const host = "127.0.0.1"

func main() {

	host := "127.0.0.1"
	port := "8080"

	/* mongo db */
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	client, _ := mongo.Connect(ctx, clientOptions()) // <--- NOT THREAD SAFE
	db.SetClint(client)

	mux := routes.NewRouter()

	/* Create the logger for the web application. */
	l := log.New()
	n := negroni.New()
	n.Use(negronilogrus.NewMiddlewareFromLogger(l, "web"))

	/*
		Create the CORS for the web application.
		this repo have no reason to use CORS middleware
		but I want to practice how to do this feature
		if you don't need CORS, just use
		n.UseHandler(mux)
	*/
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PATCH"})
	n.UseHandler(handlers.CORS(allowedOrigins, allowedMethods)(mux))

	/* Create the main server object */
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: n,
	}

	log.Println(fmt.Sprintf("Run the web server at %s:%s", host, port))
	log.Fatal(server.ListenAndServe())
}

func clientOptions() *options.ClientOptions {
	host := "db"
	if os.Getenv("profile") != "prod" {
		host = "localhost"
	}
	return options.Client().ApplyURI(
		"mongodb://" + host + ":27017",
	)
}
