package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kv4sha/golang-mongodb-docker/sources/models"
)

type sourcesService interface {
	GetSources() ([]models.Source, error)
	GetSource(id string) (*models.Source, error)
	CreateSource(source *models.Source) error
	UpdateSource(source *models.Source) error
	DeleteSource(id string) error
}

type sourcesController struct {
	sourcesService sourcesService
}

func GetSourcesController(sourcesService sourcesService) *sourcesController {
	return &sourcesController{sourcesService: sourcesService}
}

func (controller *sourcesController) GetSources(rw http.ResponseWriter, req *http.Request) {
	sources, err := controller.sourcesService.GetSources()

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonSources, _ := json.Marshal(sources)
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonSources)
}

func (controller *sourcesController) GetSource(rw http.ResponseWriter, req *http.Request) {
	sourceID := mux.Vars(req)["id"]

	source, err := controller.sourcesService.GetSource(sourceID)

	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	jsonSource, _ := json.Marshal(source)
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonSource)
}

func (controller *sourcesController) CreateSource(rw http.ResponseWriter, req *http.Request) {
	source := &models.Source{}

	if err := json.NewDecoder(req.Body).Decode(source); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := controller.sourcesService.CreateSource(source); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (controller *sourcesController) UpdateSource(rw http.ResponseWriter, req *http.Request) {
	source := &models.Source{}

	if err := json.NewDecoder(req.Body).Decode(source); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := controller.sourcesService.UpdateSource(source); err != nil {
		if err.Error() == "not found" {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (controller *sourcesController) DeleteSource(rw http.ResponseWriter, req *http.Request) {
	sourceID := mux.Vars(req)["id"]

	if err := controller.sourcesService.DeleteSource(sourceID); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
