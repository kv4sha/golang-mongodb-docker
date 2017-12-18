package main

import (
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/urfave/negroni"

	"github.com/kv4sha/golang-mongodb-docker/sources/config"
	"github.com/kv4sha/golang-mongodb-docker/sources/middlewares"
	"github.com/kv4sha/golang-mongodb-docker/sources/router"
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
