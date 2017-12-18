package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kv4sha/golang-mongodb-docker/clients/models"
)

type clientsService interface {
	GetClients() ([]models.Client, error)
	GetClient(id string) (*models.Client, error)
	CreateClient(client *models.Client) error
	UpdateClient(client *models.Client) error
	DeleteClient(id string) error
}

type clientsController struct {
	clientsService clientsService
}

func GetClientsController(clientsService clientsService) *clientsController {
	return &clientsController{
		clientsService: clientsService,
	}
}

func (controller *clientsController) GetClients(rw http.ResponseWriter, req *http.Request) {
	clients, err := controller.clientsService.GetClients()

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonClients, _ := json.Marshal(clients)
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonClients)
}

func (controller *clientsController) GetClient(rw http.ResponseWriter, req *http.Request) {
	clientID := mux.Vars(req)["id"]

	client, err := controller.clientsService.GetClient(clientID)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonClient, _ := json.Marshal(client)
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonClient)
}

func (controller *clientsController) CreateClient(rw http.ResponseWriter, req *http.Request) {
	client := &models.Client{}

	if err := json.NewDecoder(req.Body).Decode(client); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := controller.clientsService.CreateClient(client); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (controller *clientsController) UpdateClient(rw http.ResponseWriter, req *http.Request) {
	client := &models.Client{}

	if err := json.NewDecoder(req.Body).Decode(client); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := controller.clientsService.UpdateClient(client); err != nil {
		if err.Error() == "not found" {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (controller *clientsController) DeleteClient(rw http.ResponseWriter, req *http.Request) {
	clientID := mux.Vars(req)["id"]

	if err := controller.clientsService.DeleteClient(clientID); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
