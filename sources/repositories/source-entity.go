package repositories

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/kv4sha/golang-mongodb-docker/sources/models"
)

type sourceEntity struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Name       string        `bson:"name"`
	SourceType string        `bson:"sourceType"`
	Setting    interface{}   `bson:"setting"`
}

func getSourceEntity(source *models.Source) *sourceEntity {
	return &sourceEntity{
		ID:         bson.ObjectIdHex(source.ID),
		Name:       source.Name,
		Setting:    source.Setting,
		SourceType: source.SourceType,
	}
}

func (sourceEntity *sourceEntity) getSource() *models.Source {
	return &models.Source{
		ID:         sourceEntity.ID.Hex(),
		Name:       sourceEntity.Name,
		Setting:    sourceEntity.Setting,
		SourceType: sourceEntity.SourceType,
	}
}

func getSources(sourceEntities []sourceEntity) []models.Source {
	sources := []models.Source{}

	for _, sourceEntity := range sourceEntities {
		sources = append(sources, *sourceEntity.getSource())
	}

	return sources
}
