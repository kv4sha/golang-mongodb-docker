package router

import (
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"

	"github.com/kv4sha/golang-mongodb-docker/destinations/controllers"
	"github.com/kv4sha/golang-mongodb-docker/destinations/repositories"
	"github.com/kv4sha/golang-mongodb-docker/destinations/services"
)

func GetRouter(mongoDbSession *mgo.Session) *mux.Router {
	repository := repositories.GetDestinationsRepository(mongoDbSession)
	service := services.GetDestinationsService(repository)
	controller := controllers.GetDestinationsController(service)

	router := mux.NewRouter()
	router.HandleFunc("/destinations", controller.GetDestinations).Methods("GET")
	router.HandleFunc("/destinations/{id}", controller.GetDestination).Methods("GET")
	router.HandleFunc("/destinations", controller.CreateDestination).Methods("POST")
	router.HandleFunc("/destinations", controller.UpdateDestination).Methods("PUT")
	router.HandleFunc("/destinations/{id}", controller.DeleteDestination).Methods("DELETE")

	return router
}
