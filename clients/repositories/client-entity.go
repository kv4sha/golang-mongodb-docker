package repositories

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/kv4sha/golang-mongodb-docker/clients/models"
)

type clientEntity struct {
	ID             bson.ObjectId   `bson:"_id,omitempty"`
	Name           string          `bson:"name"`
	DestinationIDs []bson.ObjectId `bson:"destinationIds"`
	SourceIDs      []bson.ObjectId `bson:"sourceIds"`
}

func GetClientEntity(client *models.Client) *clientEntity {
	destinationIDs := []bson.ObjectId{}
	for _, destination := range client.Destinations {
		destinationIDs = append(destinationIDs, bson.ObjectIdHex(destination.ID))
	}

	sourceIDs := []bson.ObjectId{}
	for _, source := range client.Sources {
		sourceIDs = append(sourceIDs, bson.ObjectIdHex(source.ID))
	}

	var clientID bson.ObjectId

	if bson.IsObjectIdHex(client.ID) {
		clientID = bson.ObjectIdHex(client.ID)
	}

	return &clientEntity{
		ID:             clientID,
		Name:           client.Name,
		DestinationIDs: destinationIDs,
		SourceIDs:      sourceIDs,
	}
}

func (clientEntity *clientEntity) getClient() *models.Client {
	destinations := []models.Destination{}

	for _, destinationID := range clientEntity.DestinationIDs {
		destinations = append(destinations, models.Destination{
			ID: destinationID.Hex(),
		})
	}

	sources := []models.Source{}

	for _, sourceID := range clientEntity.SourceIDs {
		sources = append(sources, models.Source{
			ID: sourceID.Hex(),
		})
	}

	return &models.Client{
		ID:           clientEntity.ID.Hex(),
		Name:         clientEntity.Name,
		Destinations: destinations,
		Sources:      sources,
	}
}

func getClients(clientEntities []clientEntity) []models.Client {
	clients := []models.Client{}

	for _, clientEntity := range clientEntities {
		clients = append(clients, *clientEntity.getClient())
	}

	return clients
}
