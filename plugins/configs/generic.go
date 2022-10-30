package configs

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
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

// func (repo *Collection[T]) FindById(id string) (*T, error) {
// 	var target T
// 	objID, objIDerr := primitive.ObjectIDFromHex(id)
// 	if objIDerr != nil {
// 		return nil, objIDerr
// 	}
// 	err := repo.collection.FindOne(DefaultContext(), bson.M{"_id": objID}).Decode(&target)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &target, nil
// }

func (repo *Collection[T]) FindById(id string, args ...string) (T, error) {
	var target T
	objID, objIDerr := primitive.ObjectIDFromHex(id)
	if objIDerr != nil {
		return target, objIDerr
	}
	if len(args) < 1 {
		err := repo.collection.FindOne(DefaultContext(), bson.M{"_id": objID}).Decode(&target)
		if err != nil {
			return target, err
		}

		return target, nil
	} else {
		var read []T
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		condition := make([]bson.M, 0)
		condition = append(condition, bson.M{
			"$match": bson.M{
				"_id": objID,
			},
		})
		for _, val := range args {
			fnList := strings.Split(val, ".")
			field, ok := reflect.TypeOf(read).Elem().FieldByName(UcFirst(fnList[0]))
			if !ok {
				panic("Field not found")
			}

			ref := string(field.Tag.Get("ref"))
			arrayOfObjId := string(field.Tag.Get("type"))
			if arrayOfObjId != "arrayOfObjId" {
				condition = append(condition, bson.M{
					"$lookup": bson.M{
						"from":         ref,
						"localField":   val,
						"foreignField": "_id",
						"as":           val,
					}})
				if arrayOfObjId == "obj" {
					condition = append(condition, bson.M{
						"$unwind": bson.M{
							"path":                       fmt.Sprintf("$%s", val),
							"preserveNullAndEmptyArrays": true,
						}})
				}
			} else {
				asValue := fmt.Sprintf("%s_temp", fnList[0])
				condition = append(condition, bson.M{
					"$lookup": bson.M{
						"from":         ref,
						"localField":   val,
						"foreignField": "_id",
						"as":           asValue,
					}})
				condition = append(condition, bson.M{
					"$set": bson.M{
						fmt.Sprintf("%s", fnList[0]): bson.M{
							"$map": bson.M{
								"input": fmt.Sprintf("$%s", fnList[0]),
								"in": bson.M{
									"$mergeObjects": bson.A{
										"$$this",
										bson.M{
											"$arrayElemAt": bson.A{
												fmt.Sprintf("$%s", asValue), bson.M{
													"$indexOfArray": bson.A{fmt.Sprintf("$%s._id", asValue), fmt.Sprintf("$$this.%s", fnList[1])},
												},
											},
										},
									},
								},
							},
						},
					}})
			}

		}
		cursor, err := repo.collection.Aggregate(ctx, condition)
		fmt.Println("3: ", read)
		err = cursor.All(ctx, &read)
		fmt.Println("4: ", read)
		if err != nil {
			fmt.Println("err: ", err)
			return target, err
		}
		return read[0], err

	}
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
