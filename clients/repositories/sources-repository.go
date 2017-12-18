package repositories

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kv4sha/golang-mongodb-docker/clients/models"
)

type sourcesRepository struct {
	sourcesServiceURL string
}

func GetSourcesRepository(sourcesServiceURL string) *sourcesRepository {
	return &sourcesRepository{
		sourcesServiceURL: sourcesServiceURL,
	}
}

func (repository *sourcesRepository) GetByID(id string) (*models.Source, error) {
	source := &models.Source{}

	res, err := http.Get(repository.sourcesServiceURL + "/sources/" + id)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	if err := json.NewDecoder(res.Body).Decode(source); err != nil {
		return nil, err
	}

	return source, nil
}
