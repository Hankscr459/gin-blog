package configs

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection[T Document] struct {
	collection *mongo.Collection
}

type MyColl struct {
	Collection *mongo.Collection
}

func CollR[T any](collName string, dto T) *Collection[T] {
	var db = Database{}
	db.Connect(os.Getenv("MongoApplyURI"), "userdb")
	return GetCollection[T](&db, collName)
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

func (repo *Collection[T]) FindOne(query bson.M) (*T, error) {
	var target T
	err := repo.collection.FindOne(DefaultContext(), query).Decode(&target)
	if err != nil {
		return nil, err
	}
	return &target, nil
}

func CollW(collName string) *MyColl {
	var c MyColl
	var DB = MongoCN.Database("userdb")
	var Collection = DB.Collection(collName)
	c.Collection = Collection
	return &c
}

func (collW *MyColl) Create(model any) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	Id, err := collW.Collection.InsertOne(ctx, model)
	return Id, err
}

func (collW *MyColl) FindByIdAndUpdate(id string, update any) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	res := collW.Collection.FindOneAndUpdate(DefaultContext(), bson.M{"_id": objID}, bson.M{"$set": update})
	return res.Err()
}
