package configs

import (
	"context"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Document interface {
}
type Database struct {
	db     *mongo.Database
	client *mongo.Client
}

var MongoCN = connectDb()
var clientOptions = options.Client().ApplyURI(os.Getenv("MongoApplyURI"))

func connectDb() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Success to connect DB")
	return client
}

func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1
}

func (db *Database) connect(options *options.ClientOptions, dbName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
	db.client = client
	err = db.client.Ping(context.Background(), readpref.Primary())
	if err == nil {
		log.Print("Connected to MongoDB!")
	} else {
		log.Panic("Could not connect to MongoDB! Please check if mongo is running.", err)
		return err
	}
	db.db = db.client.Database(dbName)
	return nil
}

func (db *Database) Connect(connectionString string, dbName string) error {
	options := options.Client().ApplyURI(connectionString)
	err := db.connect(options, dbName)
	return err
}

func DefaultContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx
}

func GetCollection[T Document](db *Database, collectionName string) *Collection[T] {
	return &Collection[T]{db.db.Collection(collectionName)}
}
