package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kv4sha/golang-mongodb-docker/clients/services"
	"github.com/kv4sha/golang-mongodb-docker/clients/view-models"
)

type clientsService interface {
	GetClients() ([]services.Client, error)
	GetClient(id string) (*services.Client, error)

	CreateClient(client *services.Client) ([]services.Client, error)

	UpdateClient(client *services.Client) ([]services.Client, error)

	DeleteClient(id string) ([]services.Client, error)
}

type clientsController struct {
	clientsService clientsService
}

func GetClientsController(clientsService clientsService) clientsController {
	return clientsController{clientsService: clientsService}
}

func (clientsController *clientsController) GetClients(rw http.ResponseWriter, req *http.Request) {
	clients, err := clientsController.clientsService.GetClients()

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	clientsViewModel := viewModels.GetClientsViewModel(clients)

	jsonClients, _ := json.Marshal(clientsViewModel)
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonClients)
}

func (clientsController *clientsController) GetClient(rw http.ResponseWriter, req *http.Request) {
	clientID := mux.Vars(req)["id"]

	client, err := clientsController.clientsService.GetClient(clientID)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	clientViewModel := viewModels.GetClientViewModel(client)

	jsonClient, _ := json.Marshal(clientViewModel)
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonClient)
}

func (clientsController *clientsController) CreateClient(rw http.ResponseWriter, req *http.Request) {
	var clientViewModel viewModels.ClientViewModel

	if err := json.NewDecoder(req.Body).Decode(&clientViewModel); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := viewModels.GetServiceClient(&clientViewModel)

	clients, err := clientsController.clientsService.CreateClient(client)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	clientsViewModel := viewModels.GetClientsViewModel(clients)

	jsonClients, _ := json.Marshal(clientsViewModel)
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonClients)
}

func (clientsController *clientsController) UpdateClient(rw http.ResponseWriter, req *http.Request) {
	var clientViewModel viewModels.ClientViewModel

	if err := json.NewDecoder(req.Body).Decode(&clientViewModel); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := viewModels.GetServiceClient(&clientViewModel)

	clients, err := clientsController.clientsService.UpdateClient(client)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	clientsViewModel := viewModels.GetClientsViewModel(clients)

	jsonClients, _ := json.Marshal(clientsViewModel)
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonClients)
}

func (clientsController *clientsController) DeleteClient(rw http.ResponseWriter, req *http.Request) {
	clientID := mux.Vars(req)["id"]

	clients, err := clientsController.clientsService.DeleteClient(clientID)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonClients, _ := json.Marshal(clients)
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonClients)
}
