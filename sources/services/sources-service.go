package services

import "github.com/kv4sha/golang-mongodb-docker/sources/models"

type sourcesRepository interface {
	GetAll() ([]models.Source, error)
	GetByID(is string) (*models.Source, error)
	Add(source *models.Source) error
	Update(source *models.Source) error
	Delete(id string) error
}

type sourcesService struct {
	sourcesRepository sourcesRepository
}

func GetSourcesService(sourcesRepository sourcesRepository) *sourcesService {
	return &sourcesService{sourcesRepository: sourcesRepository}
}

func (service *sourcesService) GetSources() ([]models.Source, error) {
	return service.sourcesRepository.GetAll()
}

func (service *sourcesService) GetSource(id string) (*models.Source, error) {
	return service.sourcesRepository.GetByID(id)
}

func (service *sourcesService) CreateSource(source *models.Source) error {
	return service.sourcesRepository.Add(source)
}

func (service *sourcesService) UpdateSource(source *models.Source) error {
	return service.sourcesRepository.Update(source)
}

func (service *sourcesService) DeleteSource(id string) error {
	return service.sourcesRepository.Delete(id)
}
