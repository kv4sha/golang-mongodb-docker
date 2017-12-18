package repositories

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/kv4sha/golang-mongodb-docker/destinations/models"
)

type destinationsRepository struct {
	mongoDbSession *mgo.Session
}

func GetDestinationsRepository(mongoDbSession *mgo.Session) *destinationsRepository {
	return &destinationsRepository{mongoDbSession: mongoDbSession}
}

func (repository *destinationsRepository) GetAll() ([]models.Destination, error) {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	var destinationEntities []DestinationEntity

	if err := collection.Find(bson.M{}).All(&destinationEntities); err != nil {
		return nil, err
	}

	return GetDestinations(destinationEntities), nil
}

func (repository *destinationsRepository) GetByID(id string) (*models.Destination, error) {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	var destinationEntity DestinationEntity

	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("ID isn't hex")
	}

	if err := collection.FindId(bson.ObjectIdHex(id)).One(&destinationEntity); err != nil {
		return nil, err
	}

	destination := destinationEntity.GetDestination()

	return destination, nil
}

func (repository *destinationsRepository) Add(destination *models.Destination) error {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	destinationEntity := DestinationEntity{
		Name:             destination.Name,
		ConnectionString: destination.ConnectionString,
	}

	return collection.Insert(destinationEntity)
}

func (repository *destinationsRepository) Update(destination *models.Destination) error {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	destinationEntity := GetDestinationEntity(destination)

	return collection.UpdateId(bson.ObjectIdHex(destination.ID), destinationEntity)
}

func (repository *destinationsRepository) Delete(id string) error {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	if !bson.IsObjectIdHex(id) {
		return errors.New("ID isn't hex")
	}

	return collection.RemoveId(bson.ObjectIdHex(id))
}

func (repository *destinationsRepository) getCollection() *mgo.Collection {
	return repository.mongoDbSession.Copy().DB("Sourcerer").C("Destinations")
}
