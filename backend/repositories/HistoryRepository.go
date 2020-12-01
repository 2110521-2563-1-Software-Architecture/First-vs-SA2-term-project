package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/utils"
	"github.com/mitchellh/mapstructure"
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
	result, err := repo.Collection.Find(context.Background(), bson.M{"hash": key})
	if err != nil {
		return "", err
	}

	var records [10]VisitRecord
	var j = 0
	defer result.Close(context.Background())
	for result.Next(context.Background()) {
		var record bson.M
		if err = result.Decode(&record); err != nil {
			log.Fatal(err)
		}

		deResult := &VisitRecord{}
		mapstructure.Decode(record, &deResult)
		if j < 10 {
			records[j] = VisitRecord{
				Ip:        deResult.Ip,
				Hash:      deResult.Hash,
				Timestamp: deResult.Timestamp,
			}
			// fmt.Println(records)
			j++
		}
	}

	// fmt.Println("GetHistory records:", records)
	recordsMarshal, err := json.Marshal(records)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// fmt.Println(string(recordsMarshal))
	return string(recordsMarshal), err
}

func (repo *MongoHistoryRepository) CreateHistory(Ip string, Hash string, Timestamp string) (bool, error) {
	val, err := repo.Collection.InsertOne(context.Background(), bson.M{"ip": Ip, "hash": Hash, "timestamp": Timestamp})
	fmt.Println(val)
	if err != nil {
		return true, err
	}
	return true, nil
}
