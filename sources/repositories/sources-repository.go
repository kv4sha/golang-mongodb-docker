package repositories

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/kv4sha/golang-mongodb-docker/sources/models"
)

type sourcesRepository struct {
	mongoDbSession *mgo.Session
}

func GetSourcesRepository(mongoDbSession *mgo.Session) *sourcesRepository {
	return &sourcesRepository{mongoDbSession: mongoDbSession}
}

func (repository *sourcesRepository) GetAll() ([]models.Source, error) {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	var sourceEntities []sourceEntity

	if err := collection.Find(bson.M{}).All(&sourceEntities); err != nil {
		return nil, err
	}

	return getSources(sourceEntities), nil
}

func (repository *sourcesRepository) GetByID(id string) (*models.Source, error) {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	var sourceEntity sourceEntity

	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("ID isn't hex")
	}

	if err := collection.FindId(bson.ObjectIdHex(id)).One(&sourceEntity); err != nil {
		return nil, err
	}

	source := sourceEntity.getSource()

	return source, nil
}

func (repository *sourcesRepository) Add(source *models.Source) error {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	sourceEntity := sourceEntity{
		Name:       source.Name,
		SourceType: source.SourceType,

		Setting: source.Setting,
	}

	return collection.Insert(sourceEntity)
}

func (repository *sourcesRepository) Update(source *models.Source) error {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	sourceEntity := getSourceEntity(source)

	return collection.UpdateId(bson.ObjectIdHex(source.ID), sourceEntity)
}

func (repository *sourcesRepository) Delete(id string) error {
	collection := repository.getCollection()
	defer collection.Database.Session.Close()

	if !bson.IsObjectIdHex(id) {
		return errors.New("ID isn't hex")
	}

	return collection.RemoveId(bson.ObjectIdHex(id))
}

func (repository *sourcesRepository) getCollection() *mgo.Collection {
	return repository.mongoDbSession.Copy().DB("Sourcerer").C("Sources")
}
