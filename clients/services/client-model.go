package services

import (
	"github.com/kv4sha/golang-mongodb-docker/clients/repositories"
)

type Client struct {
	ID   string
	Name string
}

func GetServiceClient(repositoryClient *repositories.Client) *Client {
	return &Client{
		ID:   repositoryClient.ID,
		Name: repositoryClient.Name,
	}
}

func GetServiceClients(repositoryClients []repositories.Client) []Client {
	var serviceClients []Client

	for _, repositoryClient := range repositoryClients {
		serviceClients = append(serviceClients, Client{
			ID:   repositoryClient.ID,
			Name: repositoryClient.Name,
		})
	}

	return serviceClients
}

func GetRepositoryClient(serviceClient *Client) *repositories.Client {
	return &repositories.Client{
		ID:   serviceClient.ID,
		Name: serviceClient.Name,
	}
}

func GetRepositoryClients(serviceClients []Client) []repositories.Client {
	var repositoryClients []repositories.Client

	for _, serviceClient := range serviceClients {
		repositoryClients = append(repositoryClients, repositories.Client{
			ID:   serviceClient.ID,
			Name: serviceClient.Name,
		})
	}

	return repositoryClients
}
