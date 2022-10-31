package configs

import (
	"context"
	"gin-blog/plugins/dto"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB      = MongoCN.Database("userdb")
	UserCol = DB.Collection("users")
)

type UserService interface {
	FindByEmail(string) (*dto.ReadUserWithPassword, error)
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

func (*user) Paginate(p dto.UserPageParamsInput) (dto.UserPageParamsInput, error) {
	// filter := p.Filter
	// limit, _ := strconv.ParseInt(p.Limit, 10, 64)
	// page, _ := strconv.ParseInt(p.Page, 10, 64)
	return p, nil
}
