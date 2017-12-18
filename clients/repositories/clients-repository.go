package repositories

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/kv4sha/golang-mongodb-docker/clients/models"
)

type clientsRepository struct {
	mongoDbSession *mgo.Session
}

func GetClientsRepository(mongoDbSession *mgo.Session) *clientsRepository {
	return &clientsRepository{mongoDbSession: mongoDbSession}
}

func (repository *clientsRepository) GetAll() ([]models.Client, error) {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	var clientEntities []clientEntity

	if err := collection.Find(bson.M{}).All(&clientEntities); err != nil {
		return nil, err
	}

	return getClients(clientEntities), nil
}

func (repository *clientsRepository) GetByID(id string) (*models.Client, error) {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	var clientEntity clientEntity

	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("ID isn't hex")
	}

	if err := collection.FindId(bson.ObjectIdHex(id)).One(&clientEntity); err != nil {
		return nil, err
	}

	client := clientEntity.getClient()

	return client, nil
}

func (repository *clientsRepository) Add(client *models.Client) error {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	clientEntity := GetClientEntity(client)

	return collection.Insert(clientEntity)
}

func (repository *clientsRepository) Update(client *models.Client) error {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	clientEntity := GetClientEntity(client)

	return collection.UpdateId(bson.ObjectIdHex(client.ID), clientEntity)
}

func (repository *clientsRepository) Delete(id string) error {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	if !bson.IsObjectIdHex(id) {
		return errors.New("ID isn't hex")
	}

	return collection.RemoveId(bson.ObjectIdHex(id))
}

func (repository *clientsRepository) getCollection() *mgo.Collection {
	return repository.mongoDbSession.Copy().DB("Sourcerer").C("Clients")
}
