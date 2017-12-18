package services

import "github.com/kv4sha/golang-mongodb-docker/destinations/models"

type destinationsRepository interface {
	GetAll() ([]models.Destination, error)
	GetByID(id string) (*models.Destination, error)
	Add(destination *models.Destination) error
	Update(destination *models.Destination) error
	Delete(id string) error
}

type destinationsService struct {
	destinationsRepository destinationsRepository
}

func GetDestinationsService(destinationsRepository destinationsRepository) *destinationsService {
	return &destinationsService{destinationsRepository: destinationsRepository}
}

func (service *destinationsService) GetDestinations() ([]models.Destination, error) {
	return service.destinationsRepository.GetAll()
}

func (service *destinationsService) GetDestination(id string) (*models.Destination, error) {
	return service.destinationsRepository.GetByID(id)
}

func (service *destinationsService) CreateDestination(destination *models.Destination) error {
	return service.destinationsRepository.Add(destination)
}

func (service *destinationsService) UpdateDestination(destination *models.Destination) error {
	return service.destinationsRepository.Update(destination)
}

func (service *destinationsService) DeleteDestination(id string) error {
	return service.destinationsRepository.Delete(id)
}
