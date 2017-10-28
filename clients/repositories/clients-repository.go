package repositories

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type clientsRepository struct {
	mongoDbSession *mgo.Session
}

func GetClientsRepository(mongoDbSession *mgo.Session) *clientsRepository {
	return &clientsRepository{mongoDbSession: mongoDbSession}
}

func (repository *clientsRepository) GetAll() ([]Client, error) {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	var clients []Client

	err := collection.Find(bson.M{}).All(&clients)

	if err != nil {
		return nil, err
	}

	return clients, nil
}

func (repository *clientsRepository) GetById(id string) (*Client, error) {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	var client Client

	err := collection.Find(bson.M{"id": id}).One(&client)

	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (repository *clientsRepository) Add(client *Client) ([]Client, error) {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	err := collection.Insert(client)

	if err != nil {
		return nil, err
	}

	return repository.GetAll()
}

func (repository *clientsRepository) Update(client *Client) ([]Client, error) {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	err := collection.Update(bson.M{"id": client.ID}, client)

	if err != nil {
		return nil, err
	}

	return repository.GetAll()
}

func (repository *clientsRepository) Delete(id string) ([]Client, error) {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	err := collection.Remove(bson.M{"id": id})

	if err != nil {
		return nil, err
	}

	return repository.GetAll()
}

func (repository *clientsRepository) getCollection() *mgo.Collection {
	return repository.mongoDbSession.Copy().DB("Sourcerer").C("Clients")
}
