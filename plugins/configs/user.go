package configs

import (
	"context"
	"errors"
	"fmt"
	"gin-blog/plugins/dto"
	"reflect"
	"strings"
	"time"

	"github.com/gobeam/stringy"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB      = MongoCN.Database("userdb")
	UserCol = DB.Collection("users")
)

type Read struct {
	Name   string        `json:"name"`
	Email  string        `json:"email"`
	Friend dto.ReadUser2 `json:"friend,omitempty" bson:"friend,omitempty" ref:"users"`
}

type Read2 struct {
	Name    string         `json:"name"`
	Email   string         `json:"email"`
	Friends []dto.ReadUser `json:"friends,omitempty" bson:"friends,omitempty" ref:"users"`
}

type Read3 struct {
	Name    string      `json:"name"`
	Email   string      `json:"email"`
	Friends []ReadUser3 `json:"friends,omitempty" bson:"friends,omitempty" ref:"users"`
}

type ReadUser3 struct {
	UserId string       `json:"userId,omitempty" bson:"userId,omitempty"`
	Desc   string       `json:"desc"`
	Friend dto.ReadUser `json:"friend,omitempty" bson:"friend,omitempty"`
}

type UserService interface {
	FindByEmail(string) (*dto.ReadUserWithPassword, error)
	FindById(string) (Read, error)
	FindById2(string) (Read2, error)
	FindById3(string) (Read3, error)
	Find() ([]*dto.ReadUser, error)
}
type user struct{}

func User() UserService {
	return &user{}
}

func (*user) FindByEmail(email string) (*dto.ReadUserWithPassword, error) {
	var user *dto.ReadUserWithPassword
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	query := bson.M{"email": email}
	err := UserCol.FindOne(ctx, query).Decode(&user)
	return user, err
}

func (*user) Find() ([]*dto.ReadUser, error) {
	var list []*dto.ReadUser
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	condition := bson.M{}
	options := options.Find()
	query, err := UserCol.Find(ctx, condition, options)

	for query.Next(context.TODO()) {
		var user dto.ReadUser
		err := query.Decode(&user)
		if err != nil {
			return nil, err
		}
		list = append(list, &user)
	}
	return list, err
}

func (*user) FindById(id string) (Read, error) {
	var read []Read
	a := "ar br cr dr"
	b := strings.Split(a, " ")
	for i, v := range b {
		cv := stringy.New(v)
		fmt.Println(i, ": ", cv.UcFirst())
	}
	contains := stringy.New("friend")
	field, ok := reflect.TypeOf(read).Elem().FieldByName(contains.UcFirst())
	fmt.Println(contains.UcFirst())
	label := string(field.Tag.Get("ref"))
	if !ok {
		panic("Field not found")
	}
	fmt.Println("lable: ", label)
	objID, objIDerr := primitive.ObjectIDFromHex(id)
	if objIDerr != nil {
		return Read{}, objIDerr
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	condition := make([]bson.M, 0)
	condition = append(condition, bson.M{
		"$match": bson.M{
			"_id": objID,
		},
	})
	condition = append(condition, bson.M{
		"$lookup": bson.M{
			"from":         label,
			"localField":   "friend",
			"foreignField": "_id",
			"as":           "friend",
		}})
	condition = append(condition, bson.M{
		"$unwind": bson.M{
			"path":                       "$friend",
			"preserveNullAndEmptyArrays": true,
		}})

	cursor, err := UserCol.Aggregate(ctx, condition)
	err = cursor.All(ctx, &read)
	if err != nil {
		return Read{}, err
	}
	fmt.Println(read)
	if len(read) <= 0 {
		return Read{}, errors.New("此使用者不存在")
	}

	return read[0], nil
}

func (*user) FindById2(id string) (Read2, error) {
	var read []Read2
	objID, objIDerr := primitive.ObjectIDFromHex(id)
	if objIDerr != nil {
		return Read2{}, objIDerr
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	condition := make([]bson.M, 0)
	condition = append(condition, bson.M{
		"$match": bson.M{
			"_id": objID,
		},
	})
	condition = append(condition, bson.M{
		"$lookup": bson.M{
			"from":         "users",
			"localField":   "friends",
			"foreignField": "_id",
			"as":           "friends",
		}})
	cursor, err := UserCol.Aggregate(ctx, condition)
	err = cursor.All(ctx, &read)
	if err != nil {
		return Read2{}, err
	}
	fmt.Println(read)
	if len(read) <= 0 {
		return Read2{}, errors.New("此使用者不存在")
	}

	return read[0], nil
}

func (*user) FindById3(id string) (Read3, error) {
	var read []Read3
	objID, objIDerr := primitive.ObjectIDFromHex(id)
	if objIDerr != nil {
		return Read3{}, objIDerr
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	condition := make([]bson.M, 0)
	condition = append(condition, bson.M{
		"$match": bson.M{
			"_id": objID,
		},
	})
	condition = append(condition, bson.M{
		"$lookup": bson.M{
			"from":         "users",
			"localField":   "friends.userId",
			"foreignField": "_id",
			"as":           "fs",
		}})
	condition = append(condition, bson.M{
		"$set": bson.M{
			"friends": bson.M{
				"$map": bson.M{
					"input": "$friends",
					"in": bson.M{
						"$mergeObjects": bson.A{
							"$$this",
							bson.M{
								"friend": bson.M{
									"$arrayElemAt": bson.A{
										"$fs", bson.M{
											"$indexOfArray": bson.A{"$fs._id", "$$this.userId"},
										},
									},
								},
							},
						},
					},
				},
			},
		}})
	cursor, err := UserCol.Aggregate(ctx, condition)
	err = cursor.All(ctx, &read)
	if err != nil {
		return Read3{}, err
	}
	fmt.Println(read)
	if len(read) <= 0 {
		return Read3{}, errors.New("此使用者不存在")
	}

	return read[0], nil
}
