package configs

import (
	"context"
	"fmt"
	"gin-blog/plugins/dto"
	"strconv"
	"time"

	paginate "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB      = MongoCN.Database("userdb")
	UserCol = DB.Collection("users")
)

type UserService interface {
	FindByEmail(string) (*dto.ReadUserWithPassword, error)
	Paginate(dto.UserPageParamsInput) (dto.ReadUserPage, error)
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

func (*user) Paginate(p dto.UserPageParamsInput) (dto.ReadUserPage, error) {
	filter := bson.M{}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	limit, _ := strconv.ParseInt(p.Limit, 10, 64)
	page, _ := strconv.ParseInt(p.Page, 10, 64)
	var products []dto.ReadUser
	paginatedData, err := paginate.New(UserCol).Context(ctx).Limit(limit).Filter(filter).Page(page).Decode(&products).Find()
	if err != nil {
		fmt.Println("err: ", err)
		panic(err)
	}
	data := dto.ReadUserPage{
		Pagination: paginatedData.Pagination, Data: products,
	}
	return data, nil
}
