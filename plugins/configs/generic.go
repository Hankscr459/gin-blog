package configs

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection[T Document] struct {
	collection *mongo.Collection
}

func (repo *Collection[T]) Insert(model T) (*mongo.InsertOneResult, error) {
	Id, err := repo.collection.InsertOne(DefaultContext(), model)
	return Id, err
}

func (repo *Collection[T]) FindById(id string) (*T, error) {
	var target T
	objID, objIDerr := primitive.ObjectIDFromHex(id)
	if objIDerr != nil {
		return nil, objIDerr
	}
	err := repo.collection.FindOne(DefaultContext(), bson.M{"_id": objID}).Decode(&target)
	if err != nil {
		return nil, err
	}

	return &target, nil
}
