package viewModels

import (
	"github.com/kv4sha/golang-mongodb-docker/clients/services"
)

type ClientViewModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetClientViewModel(client *services.Client) *ClientViewModel {
	return &ClientViewModel{
		ID:   client.ID,
		Name: client.Name,
	}
}

func GetClientsViewModel(clients []services.Client) []ClientViewModel {
	var clientsViewModel []ClientViewModel

	for _, client := range clients {
		clientsViewModel = append(clientsViewModel, ClientViewModel{
			ID:   client.ID,
			Name: client.Name,
		})
	}

	return clientsViewModel
}

func GetServiceClient(clientViewModel *ClientViewModel) *services.Client {
	return &services.Client{
		ID:   clientViewModel.ID,
		Name: clientViewModel.Name,
	}
}
