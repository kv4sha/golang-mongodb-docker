package services

import "github.com/kv4sha/golang-mongodb-docker/clients/models"

type clientsRepository interface {
	GetAll() ([]models.Client, error)
	GetByID(id string) (*models.Client, error)
	Add(client *models.Client) error
	Update(client *models.Client) error
	Delete(id string) error
}

type destinationsRepository interface {
	GetByID(id string) (*models.Destination, error)
}

type sourcesRepository interface {
	GetByID(id string) (*models.Source, error)
}

type clientsService struct {
	clientsRepository      clientsRepository
	destinationsRepository destinationsRepository
	sourcesRepository      sourcesRepository
}

func GetClientsService(
	clientsRepository clientsRepository,
	destinationsRepository destinationsRepository,
	sourcesRepository sourcesRepository,
) *clientsService {
	return &clientsService{
		clientsRepository:      clientsRepository,
		destinationsRepository: destinationsRepository,
		sourcesRepository:      sourcesRepository,
	}
}

func (service *clientsService) fillDestinations(destinations []models.Destination) error {
	for destinationIndex := range destinations {
		destinationID := destinations[destinationIndex].ID

		destination, err := service.destinationsRepository.GetByID(destinationID)

		if err == nil {
			destinations[destinationIndex] = *destination
		} else {
			return err
		}
	}

	return nil
}

func (service *clientsService) fillSources(sources []models.Source) error {
	for sourceIndex := range sources {
		sourceID := sources[sourceIndex].ID

		source, err := service.sourcesRepository.GetByID(sourceID)

		if err == nil {
			sources[sourceIndex] = *source
		} else {
			return err
		}
	}

	return nil
}

func (service *clientsService) GetClients() ([]models.Client, error) {
	clients, err := service.clientsRepository.GetAll()

	if err != nil {
		return nil, err
	}

	for _, client := range clients {
		if err := service.fillDestinations(client.Destinations); err != nil {
			return clients, err
		}

		if err := service.fillSources(client.Sources); err != nil {
			return clients, err
		}
	}

	return clients, nil
}

func (service *clientsService) GetClient(id string) (*models.Client, error) {
	client, err := service.clientsRepository.GetByID(id)

	if err != nil {
		return nil, err
	}

	if err := service.fillDestinations(client.Destinations); err != nil {
		return client, err
	}

	if err := service.fillSources(client.Sources); err != nil {
		return client, err
	}

	return client, nil
}

func (service *clientsService) CreateClient(client *models.Client) error {
	return service.clientsRepository.Add(client)
}

func (service *clientsService) UpdateClient(client *models.Client) error {
	return service.clientsRepository.Update(client)
}

func (service *clientsService) DeleteClient(id string) error {
	return service.clientsRepository.Delete(id)
}
