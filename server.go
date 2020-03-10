package main

import (
	"context"
	"film36exp/db"
	"film36exp/routes"
	"fmt"
	"net/http"
	"time"

	negronilogrus "github.com/meatballhat/negroni-logrus"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"go.mongodb.org/mongo-driver/mongo"
)

// const host = "127.0.0.1"

func main() {

	host := "127.0.0.1"
	port := "8080"

	/* mongo db */
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	client, _ := mongo.Connect(ctx) // <--- NOT THREAD SAFE
	db.SetClint(client)

	mux := routes.NewRouter()

	/* Create the logger for the web application. */
	l := log.New()
	n := negroni.New()
	n.Use(negronilogrus.NewMiddlewareFromLogger(l, "web"))
	n.UseHandler(mux)

	/* Create the main server object */
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: n,
	}

	log.Println(fmt.Sprintf("Run the web server at %s:%s", host, port))
	log.Fatal(server.ListenAndServe())
}
