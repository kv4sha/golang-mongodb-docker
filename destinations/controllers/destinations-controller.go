package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kv4sha/golang-mongodb-docker/destinations/models"
)

type destinationsService interface {
	GetDestinations() ([]models.Destination, error)
	GetDestination(id string) (*models.Destination, error)
	CreateDestination(destination *models.Destination) error
	UpdateDestination(destination *models.Destination) error
	DeleteDestination(id string) error
}

type destinationsController struct {
	destinationsService destinationsService
}

func GetDestinationsController(destinationsService destinationsService) *destinationsController {
	return &destinationsController{destinationsService: destinationsService}
}

func (controller *destinationsController) GetDestinations(rw http.ResponseWriter, req *http.Request) {
	destinations, err := controller.destinationsService.GetDestinations()

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonDestinations, _ := json.Marshal(destinations)
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonDestinations)
}

func (controller *destinationsController) GetDestination(rw http.ResponseWriter, req *http.Request) {
	destinationID := mux.Vars(req)["id"]

	destination, err := controller.destinationsService.GetDestination(destinationID)

	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	jsonDestination, _ := json.Marshal(destination)
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonDestination)
}

func (controller *destinationsController) CreateDestination(rw http.ResponseWriter, req *http.Request) {
	destination := &models.Destination{}

	if err := json.NewDecoder(req.Body).Decode(destination); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := controller.destinationsService.CreateDestination(destination); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (controller *destinationsController) UpdateDestination(rw http.ResponseWriter, req *http.Request) {
	destination := &models.Destination{}

	if err := json.NewDecoder(req.Body).Decode(destination); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := controller.destinationsService.UpdateDestination(destination); err != nil {
		if err.Error() == "not found" {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (controller *destinationsController) DeleteDestination(rw http.ResponseWriter, req *http.Request) {
	destinationID := mux.Vars(req)["id"]

	if err := controller.destinationsService.DeleteDestination(destinationID); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
