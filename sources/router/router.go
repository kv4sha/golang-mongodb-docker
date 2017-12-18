package router

import (
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"

	"github.com/kv4sha/golang-mongodb-docker/sources/controllers"
	"github.com/kv4sha/golang-mongodb-docker/sources/repositories"
	"github.com/kv4sha/golang-mongodb-docker/sources/services"
)

func GetRouter(mongoDbSession *mgo.Session) *mux.Router {
	repository := repositories.GetSourcesRepository(mongoDbSession)
	service := services.GetSourcesService(repository)
	controller := controllers.GetSourcesController(service)

	router := mux.NewRouter()
	router.HandleFunc("/sources", controller.GetSources).Methods("GET")
	router.HandleFunc("/sources/{id}", controller.GetSource).Methods("GET")
	router.HandleFunc("/sources", controller.CreateSource).Methods("POST")
	router.HandleFunc("/sources", controller.UpdateSource).Methods("PUT")
	router.HandleFunc("/sources/{id}", controller.DeleteSource).Methods("DELETE")

	return router
}
