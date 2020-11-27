package repositories

import (
	"context"
	"fmt"

	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VisitRecord struct {
	Ip        string
	Hash      string
	Timestamp string
}

type HistoryRepository interface {
	GetHistory(key string) (string, error)
	CreateHistory(Ip string, Hash string, Timestamp string) (bool, error)
}

type MongoHistoryRepository struct {
	*mongo.Collection
}

func NewMongoHistoryRepository() *MongoHistoryRepository {
	uri := utils.Getenv("MONGO_URI")
	databaseName := utils.Getenv("DB_NAME")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("Mongo connection is not successful")
		panic(err)
	}
	db := client.Database(databaseName)

	fmt.Println("Mongo connection is successful")

	repo := &MongoHistoryRepository{db.Collection("history")}

	return repo
}

func (repo *MongoHistoryRepository) GetHistory(key string) (string, error) {
	result := repo.Collection.FindOne(context.Background(), bson.M{"key": key})
	var records []VisitRecord
	err := result.Decode(&records)
	if err != nil {
		return "", err
	}
	return string(""), err
}

func (repo *MongoHistoryRepository) CreateHistory(Ip string, Hash string, Timestamp string) (bool, error) {
	val, err := repo.Collection.InsertOne(context.Background(), bson.M{"ip": Ip, "hash": Hash, "timestamp": Timestamp})
	fmt.Println(val)
	if err != nil {
		return true, err
	}
	return true, nil
}
