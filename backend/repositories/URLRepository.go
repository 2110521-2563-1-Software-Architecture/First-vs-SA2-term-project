package repositories
import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/2110521-2563-1-Software-Architecture/First-vs-SA2-term-project/utils"
	"fmt"
	"errors"
)

type URLMap struct {
	Key string 
	URL string
}

type URLRepository interface {
	Create(key string, url string) (bool, error)
	GetURL(key string) (string, error)
	Exists(key string) (bool, error)
}

type MongoURLRepository struct {
	*mongo.Collection
}

func NewMongoURLRepository() *MongoURLRepository {
	uri := utils.Getenv("MONGO_URI")
	databaseName := utils.Getenv("DB_NAME")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("Mongo connection is not successful")
		panic(err)
	}
	db := client.Database(databaseName)

	fmt.Println("Mongo connection is successful")

	repo := &MongoURLRepository{db.Collection("url")}

	return repo
}

func (repo *MongoURLRepository) Exists(key string) (bool, error) {
	result := repo.Collection.FindOne(context.Background(), bson.M{"key": key})
	var k URLMap
	err := result.Decode(&k)
	if err != nil {
		// NOTE: key not found
		return false, err
	}
	return true, err
}

func (repo *MongoURLRepository) Create(key string, url string) (bool, error) {
	exists, _ := repo.Exists(key)
	if (!exists) {
		_, err := repo.Collection.InsertOne(context.Background(), bson.M{"key": key, "url": url})
		if err != nil {
			return true, err
		}

		return true, nil
	}
	return false, errors.New("Key exists")
}

func (repo *MongoURLRepository) GetURL(key string) (string, error) {
	result := repo.Collection.FindOne(context.Background(), bson.M{"key": key})
	var k URLMap
	err := result.Decode(&k)
	if err != nil {
		// NOTE: key not found
		return "", err
	}
	return k.URL, err
}