package repositories

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kv4sha/golang-mongodb-docker/clients/models"
)

type destinationsRepository struct {
	destinationsServiceURL string
}

func GetDestinationsRepository(destinationsServiceURL string) *destinationsRepository {
	return &destinationsRepository{
		destinationsServiceURL: destinationsServiceURL,
	}
}

func (repository *destinationsRepository) GetByID(id string) (*models.Destination, error) {
	destination := &models.Destination{}

	res, err := http.Get(repository.destinationsServiceURL + "/destinations/" + id)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	if err := json.NewDecoder(res.Body).Decode(destination); err != nil {
		return nil, err
	}

	return destination, nil
}
