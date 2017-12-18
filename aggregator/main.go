package main

import (
	"log"
	"net/http"

	"github.com/kv4sha/golang-mongodb-docker/aggregator/config"
	"github.com/kv4sha/golang-mongodb-docker/aggregator/middlewares"
	"github.com/kv4sha/golang-mongodb-docker/aggregator/router"

	"github.com/urfave/negroni"
)

func main() {
	appConfig := config.GetConfig()

	handler := negroni.Classic()
	handler.UseFunc(middlewares.ContentTypeFunc)
	handler.UseHandler(router.GetRouter(
		appConfig.ClientsServiceURL,
		appConfig.DestinationsServiceURL,
		appConfig.SourcesServiceURL,
	))

	server := &http.Server{
		Addr:    appConfig.Server,
		Handler: handler,
	}

	log.Println("Listening...")

	log.Fatal(server.ListenAndServe())
}
