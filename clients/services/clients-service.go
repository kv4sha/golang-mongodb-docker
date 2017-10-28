package services

import (
	"github.com/kv4sha/golang-mongodb-docker/clients/repositories"
)

type clientsRepository interface {
	GetAll() ([]repositories.Client, error)
	GetById(id string) (*repositories.Client, error)
	Add(client *repositories.Client) ([]repositories.Client, error)
	Update(client *repositories.Client) ([]repositories.Client, error)
	Delete(id string) ([]repositories.Client, error)
}

type clientsService struct {
	clientsRepository clientsRepository
}

func GetClientsService(clientsRepository clientsRepository) *clientsService {
	return &clientsService{clientsRepository: clientsRepository}
}

func (clientsService *clientsService) GetClients() ([]Client, error) {
	repositoryClients, err := clientsService.clientsRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return GetServiceClients(repositoryClients), nil
}

func (clientsService *clientsService) GetClient(id string) (*Client, error) {
	repositoryClient, err := clientsService.clientsRepository.GetById(id)

	if err != nil {
		return nil, err
	}

	return GetServiceClient(repositoryClient), nil
}

func (clientsService *clientsService) CreateClient(client *Client) ([]Client, error) {
	repositoryClient := GetRepositoryClient(client)

	repositoryClients, err := clientsService.clientsRepository.Add(repositoryClient)

	if err != nil {
		return nil, err
	}

	return GetServiceClients(repositoryClients), nil
}

func (clientsService *clientsService) UpdateClient(client *Client) ([]Client, error) {
	repositoryClient := GetRepositoryClient(client)

	repositoryClients, err := clientsService.clientsRepository.Update(repositoryClient)

	if err != nil {
		return nil, err
	}

	return GetServiceClients(repositoryClients), nil
}

func (clientsService *clientsService) DeleteClient(id string) ([]Client, error) {
	repositoryClients, err := clientsService.clientsRepository.Delete(id)

	if err != nil {
		return nil, err
	}

	return GetServiceClients(repositoryClients), nil
}
