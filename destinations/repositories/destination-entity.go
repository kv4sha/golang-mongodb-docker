package repositories

import (
	"github.com/kv4sha/golang-mongodb-docker/destinations/models"
	"gopkg.in/mgo.v2/bson"
)

type DestinationEntity struct {
	ID               bson.ObjectId `bson:"_id,omitempty"`
	Name             string        `bson:"name"`
	ConnectionString string        `bson:"connectionString"`
}

func GetDestinationEntity(destination *models.Destination) *DestinationEntity {
	return &DestinationEntity{
		ID:               bson.ObjectIdHex(destination.ID),
		Name:             destination.Name,
		ConnectionString: destination.ConnectionString,
	}
}

func (destinationEntity *DestinationEntity) GetDestination() *models.Destination {
	return &models.Destination{
		ID:               destinationEntity.ID.Hex(),
		Name:             destinationEntity.Name,
		ConnectionString: destinationEntity.ConnectionString,
	}
}

func GetDestinations(destinationEntites []DestinationEntity) []models.Destination {
	result := []models.Destination{}

	for _, destinationEntity := range destinationEntites {
		result = append(result, *destinationEntity.GetDestination())
	}

	return result
}
