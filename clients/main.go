package main

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"gopkg.in/mgo.v2"

	"github.com/kv4sha/golang-mongodb-docker/clients/config"
	"github.com/kv4sha/golang-mongodb-docker/clients/middlewares"
	"github.com/kv4sha/golang-mongodb-docker/clients/router"
)

func main() {
	appConfig := config.GetConfig()

	mongoDbSession, err := mgo.Dial(appConfig.MongoDbURL)
	defer mongoDbSession.Close()

	if err != nil {
		log.Fatal("[MondoDb]: ", err)
	}

	handler := negroni.Classic()
	handler.UseFunc(middlewares.ContentTypeFunc)
	handler.UseHandler(router.GetRouter(mongoDbSession))

	server := &http.Server{
		Addr:    appConfig.Server,
		Handler: handler,
	}

	log.Println("Listening...")

	log.Fatal(server.ListenAndServe())
}
