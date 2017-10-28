package router

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	"github.com/kv4sha/golang-mongodb-docker/clients/controllers"
	"github.com/kv4sha/golang-mongodb-docker/clients/repositories"
	"github.com/kv4sha/golang-mongodb-docker/clients/services"
)

func GetRouter(mongoDbSession *mgo.Session) *mux.Router {
	repository := repositories.GetClientsRepository(mongoDbSession)

	service := services.GetClientsService(repository)

	controller := controllers.GetClientsController(service)

	router := mux.NewRouter()
	router.HandleFunc("/clients", controller.GetClients).Methods("GET")
	router.HandleFunc("/clients/{id}", controller.GetClient).Methods("GET")
	router.HandleFunc("/clients", controller.CreateClient).Methods("POST")
	router.HandleFunc("/clients", controller.UpdateClient).Methods("PUT")
	router.HandleFunc("/clients/{id}", controller.DeleteClient).Methods("DELETE")

	return router
}
