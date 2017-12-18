package router

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	"github.com/kv4sha/golang-mongodb-docker/clients/controllers"
	"github.com/kv4sha/golang-mongodb-docker/clients/repositories"
	"github.com/kv4sha/golang-mongodb-docker/clients/services"
)

func GetRouter(mongoDbSession *mgo.Session, destinationsServiceURL string, sourcesServiceURL string) *mux.Router {
	clientsRepository := repositories.GetClientsRepository(mongoDbSession)
	destinationsRepository := repositories.GetDestinationsRepository(destinationsServiceURL)
	sourcesRepository := repositories.GetSourcesRepository(sourcesServiceURL)

	service := services.GetClientsService(clientsRepository, destinationsRepository, sourcesRepository)
	controller := controllers.GetClientsController(service)

	router := mux.NewRouter()
	router.HandleFunc("/clients", controller.GetClients).Methods("GET")
	router.HandleFunc("/clients/{id}", controller.GetClient).Methods("GET")
	router.HandleFunc("/clients", controller.CreateClient).Methods("POST")
	router.HandleFunc("/clients", controller.UpdateClient).Methods("PUT")
	router.HandleFunc("/clients/{id}", controller.DeleteClient).Methods("DELETE")

	return router
}
